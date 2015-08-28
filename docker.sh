#!/bin/bash

build() {
  cd postgres;
  docker build --no-cache -t outyet1 .
  boot2docker up && $(boot2docker shellinit) 
  boot2docker ip
  cd ../webserver;
  docker build --no-cache -t outyet2 .
  boot2docker up && $(boot2docker shellinit) 
  boot2docker ip
}

test() {
  # start docker postgresql_db
  docker run -d -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=changeme --name postgresql_db -v "/var/lib/postgresql/data:/var/lib/postgresql/data" outyet1
  # run test
  docker run -ti --rm outyet2 go test -v ./authentication
  # stop docker postgresql_db
  docker kill postgresql_db
  docker rm postgresql_db
}

run() {
  # Run a docker
  # -v ... it maps the filesystem from container to docker host filesystem
  # the difference between docker run -v /host/path:/container/path and docker run -v /container/path is, 
  # that in the first case you provide a directory to mount in, in the 2nd case, docker creates a directory 
  # in /var/lib/docker and uses that as the host path
  docker run -d -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=changeme --name postgresql_db -v "/var/lib/postgresql/data:/var/lib/postgresql/data" outyet1
  docker run -d -p 8080:8080 --name webserver -v "/var/volume1:/go/src/github.com/geraldstanje/web_app/webserver/files" outyet2
}

info() {
  # docker ps
  # docker inspect hash
  docker ps
}

logs() {
  docker logs postgresql_db
  docker logs webserver
}

stop() {
  # docker stop with name
  # stop docker postgresql_db
  docker kill postgresql_db
  docker rm postgresql_db

  # stop docker webserver
  docker kill webserver
  docker rm webserver
}

format() {
  boot2docker ssh 'sudo rm -rf /var/lib/postgresql; sudo rm -rf /var/volume1'
}

# Delete all containers
# docker rm $(docker ps -a -q)

# Delete all images
# docker rmi $(docker images -q)

case $1 in build|test|run|info|logs|stop|format) "$1" ;; *) printf >&2 '%s: unknown command\n' "$1"; exit 1;; esac