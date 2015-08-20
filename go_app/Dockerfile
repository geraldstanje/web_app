FROM golang

RUN ["go", "get", "github.com/geraldstanje/web_app/go_app" ]

WORKDIR /go/src/github.com/geraldstanje/web_app/go_app

ENTRYPOINT go build && ./main

EXPOSE 8080