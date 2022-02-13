FROM ubuntu:latest

RUN apt-get update

ADD https://github.com/pressly/goose/releases/download/v2.6.0/goose-linux64 /bin/goose
RUN chmod +x /bin/goose

RUN mkdir -p /migrations
WORKDIR /migrations
COPY ./migrations .
RUN chmod +x entrypoint.sh

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

CMD /wait && ./entrypoint.sh