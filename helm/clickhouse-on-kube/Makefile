
helm-dry:
	helm install --name clickhouse-test --namespace=playground --debug --dry-run .

helm-lint:
	helm lint .

helm-install-test:
	helm install --name clickhouse-test --namespace=playground .

helm-install-stg:
	helm install --name clickhouse-stg \
	    --namespace=clickhouse \
	    --set replicaCount=5 \
	    --set ebs.size=500Gi \
	    --set security.enabled=true \
	    --set resources.limits.cpu=3800m \
	    --set resources.limits.memory=28000Mi \
	    --set resources.requests.cpu=3800m \
	    --set resources.requests.memory=28000Mi \
	    .

helm-del:
	helm del --purge clickhouse-test

kube-get:
	kubectl -n playground get deployment,pod,svc,ingress
