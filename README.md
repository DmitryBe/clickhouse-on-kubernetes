# clickhouse-on-kubernetes

Run clickhouse server on kubernetes

## Quick start

```bash
# clone repo
git clone https://github.com/DmitryBe/clickhouse-on-kubernetes.git
cd clickhouse-on-kubernetes/helm/clickhouse-on-kube

# update values.yaml
# replicaCount: 3 # number of nodes && shards (1 replica)
# zk.servers: <your zk address> # or disable
# nodeSelector.type: <node label> # or remove
# disable elb (if running on google cloud | minikube)
# resources: ...

# install chart
helm install \
    --name clickhouse-test \
    --namespace=default .

# run clickhouse-client
docker run -it --rm --net=host yandex/clickhouse-client -h <elb_endpoint>

# check cluster config
:) select * from system.clusters

# create db and distrib table
CREATE DATABASE IF NOT EXISTS DB01;
dd
USE DB01;

CREATE TABLE IF NOT EXISTS DB01.Table1 ON CLUSTER default_cluster
(
    UpdateDate Date,    
    GeoHash String,    
    NBooked UInt32  
) ENGINE = MergeTree(UpdateDate, (GeoHash, UpdateDate), 8192);

CREATE TABLE IF NOT EXISTS DB01.Table1_c ON CLUSTER default_cluster as DB01.Table1
ENGINE = Distributed(default_cluster, DB01, Table1, cityHash64(GeoHash));

# insert some data
INSERT INTO DB01.Table1_c VALUES ('2017-11-08', 'geo01', 1);
INSERT INTO DB01.Table1_c VALUES ('2017-11-08', 'geo02', 2);
INSERT INTO DB01.Table1_c VALUES ('2017-11-08', 'geo03', 3);

# query distrib table
SELECT * FROM DB01.Table1_c
```

## Dev

To rebuild docker image run: `make docker-build`

To rebuild helm package run: `make helm-pckg`
