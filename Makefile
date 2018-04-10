REPO=dmitryb/clickhouse-server
TAG=latest

# build clickhouse-server docker image
docker-build:
	docker build -t $(REPO):$(TAG) .

# push clickhouse-server docker image to hub
docker-push:
	docker push $(REPO):$(TAG)

# create helm package and reindex
helm-pckg:
	helm package ./helm/clickhouse-on-kube
	helm repo index .
