FROM golang:1.17 AS builder

RUN go version

COPY . /github.com/1makarov/go-crud-example/
WORKDIR /github.com/1makarov/go-crud-example/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/1makarov/go-crud-example/.bin/app .
COPY --from=0 /github.com/1makarov/go-crud-example/configs configs/

CMD ["./app"]