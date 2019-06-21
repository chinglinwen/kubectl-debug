export GO111MODULE=on
go build -o kubectl-debug
curl fs.devops.haodai.net/k8s/uploadapi -F file=@kubectl-debug -F truncate=yes
