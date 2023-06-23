FROM golang:1.15

WORKDIR /go/src/github.com/redis_docker

COPY . .

RUN go get -u github.com/go-redis/redis

RUN go build -o main .

CMD ["./main"] 
 
