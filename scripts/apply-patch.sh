# requried env
# KUBE_NAMESPACE
# patch.json

KUBE_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
if [ -z "$KUBE_TOKEN" ]; then
  echo "Error: KUBE_TOKEN is empty"
  exit 1
fi

if [ -z "$KUBE_NAMESPACE" ]; then
  echo "Error: KUBE_NAMESPACE is empty"
  exit 1
fi

HOSTNAME=$(hostname)

api_url="https://$KUBERNETES_SERVICE_HOST:$KUBERNETES_PORT_443_TCP_PORT/api/v1/namespaces/$KUBE_NAMESPACE/pods/$HOSTNAME"

echo "Applying JSON patch:"
cat patch.json
echo "at $api_url"

curl -sSk --request PATCH --data "$(cat patch.json)" \
  -H "Authorization: Bearer $KUBE_TOKEN" \
  -H "Content-Type:application/json-patch+json" \
  "$api_url"

