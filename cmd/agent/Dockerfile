FROM harbor.haodai.net/base/alpine:3.7cgo
WORKDIR /mnt/container

MAINTAINER wenzhenglin(http://g.haodai.net/wenzhenglin/debug-agent)

ADD ./agent /app/debug-agent

CMD /app/debug-agent
ENTRYPOINT ["/app/debug-agent"]

EXPOSE 10027
