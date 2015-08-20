#!/bin/bash

build() {
  docker build --no-cache -t outyet .
  boot2docker up && $(boot2docker shellinit) 
  boot2docker ip
}

run() {
  # Run a docker
  # -v ... it maps the filesystem from container to docker host filesystem
  docker run -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=changeme --name postgresql_example_name -e 'DB_USER=dbuser' -e 'DB_PASS=dbpass' -e 'DB_NAME=dbname' -v "/var/lib/postgresql/data:/var/lib/postgresql/data" outyet
}

info() {
  # docker ps
  docker ps
  # docker inspect hash
}

stopall() {
  # docker stop hash
  docker stop $(docker ps -a -q)
}

cleanall() {
  # Remove all containers
  docker rm $(docker ps -a -q)
  # Romove all images
  #docker rmi $(docker images -q)
}

case $1 in build|run|info|stopall|cleanall) "$1" ;; *) printf >&2 '%s: unknown command\n' "$1"; exit 1;; esac