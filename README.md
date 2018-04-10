# clickhouse-on-kubernetes

## misc

clickhouse remote-servers config path: /etc/clickhouse-server/config.d/clickhouse_remote_servers.xml

example: 
``xml
<yandex>
    <remote_servers incl="clickhouse_remote_servers" >        
    <cluster_3shards_1replicas>
        
    <shard>
        <replica>
            <host>172.17.0.3</host>
            <port>9000</port>
        </replica>
    </shard>    
    
    </cluster_3shards_1replicas>
    </remote_servers>
</yandex>
```
