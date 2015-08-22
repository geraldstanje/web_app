FROM golang

RUN ["go", "get", "github.com/geraldstanje/web_app/webserver" ]

WORKDIR /go/src/github.com/geraldstanje/web_app/webserver

#VOLUME ["/var/volume1", "/go/src/github.com/geraldstanje/web_app/webserver/files"]

ENTRYPOINT go build && ./webserver

EXPOSE 8080