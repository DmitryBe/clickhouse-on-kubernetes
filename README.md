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
# resources: ...

# install chart
helm install \
    --name clickhouse-test \
    --namespace=default .
```

## Dev

To rebuild docker image run: `make docker-build`

To rebuild helm package run: `make helm-pckg`
