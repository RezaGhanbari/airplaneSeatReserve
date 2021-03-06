FROM golang:latest

RUN mkdir /app

ADD . /app/
WORKDIR /app


RUN go get github.com/gin-gonic/gin
RUN go get github.com/satori/go.uuid

EXPOSE 8000:8000
RUN go build -o main .
CMD ["/app/main"]
