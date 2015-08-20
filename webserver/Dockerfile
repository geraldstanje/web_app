FROM golang

RUN ["go", "get", "github.com/geraldstanje/web_app/webserver" ]

WORKDIR /go/src/github.com/geraldstanje/web_app/webserver

ENTRYPOINT go build && ./webserver

EXPOSE 8080