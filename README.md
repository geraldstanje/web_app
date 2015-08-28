# Music Album Collections:

## Install and Run:

### Build docker images
This builds the webserver and postgres docker image <br />
$ ./docker.sh build

### Run docker container tests
This command calls go test to test all packages <br />
$ ./docker.sh test

### Run docker containers
This starts the webserver and postgres docker container <br />
$ ./docker.sh run

### Display logs:
This shows the all logs within the webserver and postgres docker container <br />
$ ./docker.sh logs

### Stop docker containers
$ ./docker.sh stop

### Delete docker container volumes
Delete the volume of the postgres db <br />
Delete the volume of the image directory <br />
$ ./docker.sh format

## Next, Next Steps:
1. Add password authentication strategies: store SHA512 or bcrypt instead of the plain password
2. Create a database table for images, and only keep the url of the image, change postgres schema: 
  * CREATE TABLE users (id UUID PRIMARY KEY, email CITEXT(200) UNIQUE, password TEXT NOT NULL);
  * CREATE TABLE images (id UUID PRIMARY KEY, user_id UUID REFERENCES users, url TEXT);
3. Store/upload the images to a CDN server
4. Add Search feature (the user should be able to search for his own albums) 
4. Improve error handling in webserver and postgres
5. Add testcases for each package

---

#### Postgres Infos:

##### Get boot2docker IP
$ boot2docker ip <br />
 The VM's Host only interface IP address is: 192.168.59.103

##### Start PostgreSQL interactive terminal
$ psql -h 192.168.59.103 -p 5432 -d docker -U admin --password