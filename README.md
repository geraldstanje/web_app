# Music Album Collections:

## Install and Run:

### Build docker images
This builds the webserver and postgres docker image <br />
$ ./docker.sh build

### Run docker containers
This starts the webserver and postgres docker container <br />
$ ./docker.sh run

### Display logs:
This shows the all logs within the webserver and postgres docker container <br />
$ ./docker.sh logs

### Stop docker containers
$ ./docker.sh stopall

### Remove docker containers
$ ./docker.sh cleanall

## Postgres Infos:

### Get boot2docker IP
$ boot2docker ip <br />
 The VM's Host only interface IP address is: 192.168.59.103

### Start PostgreSQL interactive terminal
$ psql -h 192.168.59.103 -p 5432 -d docker -U admin --password

### Delete the volume for the postgres db
$ boot2docker ssh <br />
$ sudo rm -rf /var/lib/postgresql

### Delete the volume for the image directory
$ boot2docker ssh <br />
$ sudo rm -rf /var/volume1

## Next, Next Steps:
1. Add password authentication strategies: store SHA512 or bcrypt instead of the plain password
2. create a database table for images, change postgres schema: 
  * CREATE TABLE users (id UUID PRIMARY KEY, email CITEXT(200) UNIQUE, password TEXT NOT NULL);
  * CREATE TABLE images (id UUID PRIMARY KEY, user_id UUID REFERENCES users, url TEXT);
3. Store/upload the images to a CDN server
4. Add better error handling in webserver and postgres
5. Add testcases