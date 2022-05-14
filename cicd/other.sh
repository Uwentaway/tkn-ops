kubectl create secret docker-registry dockerhub --docker-server=https://index.docker.io/v1/ --docker-username=gitopsliu --docker-password=qweASD123 --dry-run=client -o json | jq -r '.data.".dockerconfigjson"' | base64 -d > /tmp/config.json && kubectl create secret generic docker-config -n demo  --from-file=/tmp/config.json && rm -f /tmp/config.json

kubectl create secret generic kubernetes-config -n demo  --from-file=/root/.kube/config
