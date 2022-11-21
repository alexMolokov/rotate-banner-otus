FROM ubuntu:latest

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

ADD https://github.com/pressly/goose/releases/download/v2.6.0/goose-linux64 /bin/goose
RUN chmod +x /bin/goose

RUN mkdir -p /migrations
WORKDIR /migrations
COPY ./migrations .
ADD ./deployments/wait-for-postgres.sh .
RUN chmod +x wait-for-postgres.sh
RUN chmod +x entrypoint.sh