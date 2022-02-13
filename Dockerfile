FROM golang:1.16-stretch

ENV BIN_FILE /usr/src/app/bin/rotator

RUN mkdir -p /usr/src/app/log
WORKDIR /usr/src/app
COPY . /usr/src/app

# make build
RUN make build ; chmod a+x ./bin/rotator
ENV TZ Europe/Moscow

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

ARG CONFIG_FILE_NAME

ENV CONFIG_FILE /usr/src/app/bin/${CONFIG_FILE_NAME}.json
COPY ./configs/${CONFIG_FILE_NAME}.json ${CONFIG_FILE}

CMD /wait && ${BIN_FILE} -config ${CONFIG_FILE}