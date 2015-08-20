#!/bin/bash

build() {
  docker build --no-cache -t outyet .
  boot2docker up && $(boot2docker shellinit) 
  boot2docker ip
}

run() {
  # Run a docker
  docker run -p 8080:8080 -t outyet -name webserver
}

info() {
  # docker ps
  docker ps
  # docker inspect hash
}

stop() {
  # docker stop with name
  docker stop webserver
}

stopall() {
  # docker stop hash
  docker stop $(docker ps -a -q)
}

clean() {
  # Remove docker with name
  docker rm webserver
}

cleanall() {
  # Remove all containers
  docker rm $(docker ps -a -q)
  # Romove all images
  #docker rmi $(docker images -q)
}

case $1 in build|run|info|stop|stopall|clean|cleanall) "$1" ;; *) printf >&2 '%s: unknown command\n' "$1"; exit 1;; esac