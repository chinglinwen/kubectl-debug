set -e
# export GO111MODULE=on
go build -o gobash
curl fs.devops.haodai.net/soft/uploadapi -F file=@gobash -F truncate=yes
cksum gobash

