FROM golang:1.10.0-alpine3.7

WORKDIR /go/src/github.com/azeezolaniran2016/go-chat

RUN apk add --no-cache make git bash curl gcc musl-dev && \ 
    go get github.com/codegangsta/gin && \
    go get -u github.com/kardianos/govendor

CMD gin -i -a 80 run

