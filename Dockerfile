FROM golang:1.16.2-alpine

WORKDIR /go/app

RUN go mod init go-word-reference

RUN go get -u github.com/cosmtrek/air

CMD [ "air" ]