# Install Postgres DB:

## Build docker image
$ ./docker build

## Run docker image
$ ./docker run

## Get boot2docker IP
$ boot2docker ip <br />
 The VM's Host only interface IP address is: 192.168.59.103

## Start PostgreSQL interactive terminal
$ psql -h localhost -p 5432 -d docker -U admin --password