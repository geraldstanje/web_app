FROM golang

RUN ["go", "get", "github.com/geraldstanje/web_app/webserver" ]

WORKDIR /go/src/github.com/geraldstanje/web_app/webserver

VOLUME ["/var/volume1", "/var/volume2"]

ENTRYPOINT go build && ./webserver

EXPOSE 8080