
helm-dry:
	helm install --name clickhouse-test --namespace=playground --debug --dry-run .

helm-lint:
	helm lint .

helm-install:
	helm install --name clickhouse-test --namespace=playground .

helm-del:
	helm del --purge clickhouse-test

kube-get:
	kubectl -n playground get deployment,pod,svc,ingress
