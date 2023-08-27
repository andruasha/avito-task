FROM golang:latest

WORKDIR /app

COPY . .

RUN go get -d -v github.com/go-redis/redis/v8

RUN go build -o app

CMD ["./app"]
