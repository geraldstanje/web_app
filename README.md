# Music Album Collections Project:

## Install and Run:

### Build docker images
This builds the webserver and postgres db
$ ./docker.sh build

### Run docker images
This starts the webserver and postgres db
$ ./docker.sh run

### Display logs:
This shows the entire log
$ ./docker.sh logs

### Stop docker images
$ ./docker.sh stopall

### Clean docker images
$ ./docker.sh cleanall

## Postgres Infos:

### Get boot2docker IP
$ boot2docker ip <br />
 The VM's Host only interface IP address is: 192.168.59.103

### Start PostgreSQL interactive terminal
$ psql -h 192.168.59.103 -p 5432 -d docker -U admin --password

### Delete the volume for the postgres db
$ boot2docker ssh
$ sudo rm -rf /var/lib/postgresql

### Delete the volume for the image directory
$ boot2docker ssh
$ sudo rm -rf /var/volume1