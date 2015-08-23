FROM golang

RUN ["go", "get", "github.com/geraldstanje/web_app/webserver" ]

WORKDIR /go/src/github.com/geraldstanje/web_app/webserver

# Add directories and files
ADD templates/index.html /go/src/github.com/geraldstanje/web_app/webserver/templates/index.html
ADD templates/musicalbums.html /go/src/github.com/geraldstanje/web_app/webserver/templates/musicalbums.html

ENTRYPOINT go build && ./webserver

EXPOSE 8080