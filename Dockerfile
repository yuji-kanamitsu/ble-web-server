FROM golang:1.14

RUN mkdir /go/src/docker-test

WORKDIR /go/src/docker-test

RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/mattn/go-sqlite3

ADD . /go/src/docker-test

CMD ["go", "run", "main.go"]