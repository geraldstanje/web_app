FROM golang

RUN ["go", "get", "github.com/geraldstanje/web_app/webserver" ]

WORKDIR /go/src/github.com/geraldstanje/web_app/webserver

# Add directories and files
ADD static/index.html /go/src/github.com/geraldstanje/web_app/webserver/static/index.html
ADD templates/musicalbums.html /go/src/github.com/geraldstanje/web_app/webserver/templates/musicalbums.html

# ENTRYPOINT go build && ./webserver

RUN go build
ENTRYPOINT ./webserver

EXPOSE 8080