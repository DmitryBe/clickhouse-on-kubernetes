REPO=grabds/clickhouse-server
TAG=latest

docker-build:
	docker build -t $(REPO):$(TAG) .

docker-push:
	docker push $(REPO):$(TAG)
