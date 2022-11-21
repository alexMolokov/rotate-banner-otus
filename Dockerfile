FROM golang:1.16-stretch

RUN mkdir -p /usr/src/app/log
WORKDIR /usr/src/app
COPY . /usr/src/app

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

ADD ./deployments/wait-for-postgres.sh .
# make wait-for-postgres.sh executable
RUN chmod a+x wait-for-postgres.sh

# make buildd
RUN make build ; chmod a+x ./bin/rotator

ENV TZ Europe/Moscow