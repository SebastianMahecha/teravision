FROM golang:latest

RUN mkdir -p $GOPATH/src/github.com/teravision
WORKDIR $GOPATH/src/github.com/teravision

ADD . .
RUN go get -u -v github.com/gin-gonic/gin
RUN go get -u -v github.com/jinzhu/gorm
RUN go get -u -v github.com/jinzhu/gorm/dialects/postgres
RUN go get -u -v github.com/lib/pq

RUN go build ./application.go

EXPOSE 3000

CMD ["./application"]

