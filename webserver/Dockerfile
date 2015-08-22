FROM golang

RUN ["go", "get", "github.com/geraldstanje/web_app/webserver" ]

WORKDIR /go/src/github.com/geraldstanje/web_app/webserver

ADD files /go/src/github.com/geraldstanje/web_app/

ENTRYPOINT go build && ./webserver

EXPOSE 8080