FROM golang

RUN ["go", "get", "github.com/geraldstanje/web_app/go_app" ]

WORKDIR /go/src/github.com/geraldstanje/web_app/go_app

ENTRYPOINT go build && ./go_app

EXPOSE 8080