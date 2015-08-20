#!/bin/bash

build() {
  docker build --no-cache -t outyet .
  boot2docker up && $(boot2docker shellinit) 
  boot2docker ip
}

run() {
  # Run a docker
  # -v ... it maps the filesystem from container to docker host filesystem
  docker run -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=changeme -name postgresql_db -v "/var/lib/postgresql/data:/var/lib/postgresql/data" outyet
}

info() {
  # docker ps
  docker ps
  # docker inspect hash
}

stop() {
  # docker stop with name
  docker stop postgresql_db
}

stopall() {
  # docker stop hash
  docker stop $(docker ps -a -q)
}

clean() {
  # Remove container with name
  docker rm postgresql_db
}

cleanall() {
  # Remove all containers
  docker rm $(docker ps -a -q)
  # Romove all images
  #docker rmi $(docker images -q)
}

case $1 in build|run|info|stop|stopall|clean|cleanall) "$1" ;; *) printf >&2 '%s: unknown command\n' "$1"; exit 1;; esac