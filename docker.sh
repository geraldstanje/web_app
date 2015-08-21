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

run() {
  # Run a docker
  # -v ... it maps the filesystem from container to docker host filesystem
  cd postgres;
  docker run -d -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=changeme --name postgresql_db -v "/var/lib/postgresql/data:/var/lib/postgresql/data" outyet1
  cd ../webserver;
  docker run -d -p 8080:8080 --name webserver outyet2
}

info() {
  # docker ps
  # docker inspect hash
  cd postgres;
  docker ps
  cd ../webserver;
  docker ps;
}

logs() {
  docker logs postgresql_db
  docker logs webserver
}

stop() {
  # docker stop with name
  cd postgres;
  docker stop postgresql_db
  cd ../webserver;
  docker stop webserver
}

stopall() {
  # docker stop hash
  docker stop $(docker ps -a -q)
}

clean() {
  # Remove container with name
  cd postgres;
  docker rm postgresql_db
  cd ../webserver;
  docker stop webserver
}

cleanall() {
  # Remove all containers
  docker rm $(docker ps -a -q)
  # Romove all images
  #docker rmi $(docker images -q)
}

case $1 in build|run|info|logs|stop|stopall|clean|cleanall) "$1" ;; *) printf >&2 '%s: unknown command\n' "$1"; exit 1;; esac