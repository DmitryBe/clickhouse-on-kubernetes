package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// RemoteServersTemplate provides template for cluster remote server
const RemoteServersTemplate = `
<yandex>
	<remote_servers incl="clickhouse_remote_servers" >
		<%s>
			%s
		</%s>
	</remote_servers>
</yandex>
`

// ShardTemplate templates for node
const ShardTemplate = `
			<shard>
				<replica>
					<host>%s</host>
					<port>%s</port>
				</replica>
			</shard>
`

// UpdateClusterConfig is blah
func UpdateClusterConfig(nodePorts []string, clusterName string) string {
	fmt.Println("updating cluster config")
	shards := []string{}
	for _, nodePort := range nodePorts {
		parts := strings.Split(nodePort, ":")
		node, port := parts[0], parts[1]
		shard := fmt.Sprintf(ShardTemplate, node, port)
		shards = append(shards, shard)
	}
	return fmt.Sprintf(RemoteServersTemplate, clusterName, strings.Join(shards, ""), clusterName)
}

// WriteStringToFile is
func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}
