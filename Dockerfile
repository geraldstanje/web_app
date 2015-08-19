FROM postgres:latest

RUN locale-gen en_US.UTF-8
RUN update-locale LANG=en_US.UTF-8

COPY ./init-scripts/* /docker-entrypoint-initdb.d/