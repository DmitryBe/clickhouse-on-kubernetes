package main

import (
	"fmt"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// ClickHousePort is a default clickhous port
const ClickHousePort = 9000

// ClickHouseConfigLocation config path
const ClickHouseConfigD = "/etc/clickhouse-server/config.d"

// ClickhouseRemoteServersConfigFileName - cluster config file name
const ClickhouseRemoteServersConfigFileName = "clickhouse_remote_servers.xml"

// ClickHouseClusterName for remote servres config
const ClickHouseClusterName = "default_cluster"

func main() {
	fmt.Println("starting coordinator")

	// get zk connection
	zksStr := os.Getenv("ZOOKEEPER_SERVERS")
	if zksStr == "" {
		panic("ZOOKEEPER_SERVERS is requried")
	}
	fmt.Println("zk: " + zksStr)

	// get zk key path
	zkPath := os.Getenv("ZOOKEEPER_PATH")
	if zkPath == "" {
		zkPath = "/clickhouse-coordinator-01"
	}
	fmt.Println("zk-path: " + zkPath)

	clickHouseConfigD := os.Getenv("CLICKHOUSE_CONFIG_LOCATION")
	if clickHouseConfigD == "" {
		clickHouseConfigD = ClickHouseConfigD
	}
	fmt.Println("clickhouse config.d path: " + clickHouseConfigD)

	// connect to zk
	conn := zkConnect()
	defer conn.Close()

	// create root path
	path, err := conn.Create(zkPath, []byte(""), int32(0), zk.WorldACL(zk.PermAll))
	if err != nil {
		fmt.Printf("zp path exists: %+v\n", zkPath)
	}
	fmt.Printf("create: %+v\n", path)

	// start zk watcher
	snapshots, errors := zkMirror(conn, zkPath)
	// watching events
	go func() {
		for {
			select {
			case snapshot := <-snapshots:
				remoteServerConfig := UpdateClusterConfig(snapshot, ClickHouseClusterName)
				fmt.Println("--------------------------------------------------------")
				fmt.Println(remoteServerConfig)
				fmt.Println("--------------------------------------------------------")
				if len(snapshot) > 0 {
					// update config file
					remoteServreConfigPath := fmt.Sprintf("%s/%s", clickHouseConfigD, ClickhouseRemoteServersConfigFileName)
					fmt.Println("updating remote server config: " + remoteServreConfigPath)
					err := WriteStringToFile(remoteServreConfigPath, remoteServerConfig)
					must(err)
				}

			case err := <-errors:
				panic(err)
			}
		}
	}()

	// worming up
	time.Sleep(time.Second)

	// get pod ip
	ip := GetLocalIP()
	fmt.Println(ip.String())

	// create ephemeral path
	flags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)
	clickhouseHostPort := fmt.Sprintf("%s:%v", ip, ClickHousePort)
	zkNodePath := fmt.Sprintf("%s/%s", zkPath, clickhouseHostPort)
	fmt.Println("create zk path for node: " + zkNodePath)
	_, err = conn.Create(zkNodePath, []byte(clickhouseHostPort), flags, acl)
	must(err)

	// wait for termination
	fmt.Println("waiting for termination")
	_, _, ech, err := conn.ExistsW(zkPath)
	must(err)
	evt := <-ech
	fmt.Println("closing coordinator")
	must(evt.Err)
}
