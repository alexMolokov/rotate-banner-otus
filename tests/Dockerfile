FROM golang:1.16-stretch

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . /usr/src/app

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

CMD /wait && go test ./tests/...