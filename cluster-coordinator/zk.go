package main

import (
	"log"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func zkConnect() *zk.Conn {
	zksStr := "0.0.0.0"
	zks := strings.Split(zksStr, ",")
	conn, _, err := zk.Connect(zks, time.Second)
	must(err)
	return conn
}

func zkMirror(conn *zk.Conn, path string) (chan []string, chan error) {
	snapshots := make(chan []string)
	errors := make(chan error)
	go func() {
		for {
			snapshot, _, events, err := conn.ChildrenW(path)
			if err != nil {
				errors <- err
				return
			}

			// snaps pipe
			snapshots <- snapshot

			// error pipe
			evt := <-events
			if evt.Err != nil {
				errors <- evt.Err
				return
			}

			// rest
			time.Sleep(time.Second * 3)
		}
	}()
	return snapshots, errors
}
