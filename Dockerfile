FROM golang:latest

RUN mkdir /app

ADD . /app/
WORKDIR /app

ENV PORT=8000
EXPOSE $PORT:$PORT

RUN go get github.com/gin-gonic/gin
RUN go build -o main .
CMD ["/app/main"]
