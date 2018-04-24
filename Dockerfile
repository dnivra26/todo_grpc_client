FROM golang:alpine AS builder
RUN apk add --no-cache --update git
ENV GOPATH=/go

ADD . $GOPATH/src/todo_grpc_client
WORKDIR $GOPATH/src/todo_grpc_client

RUN go get

RUN go build -v -o /go/bin/todo_grpc_client server.go

FROM alpine:latest

COPY --from=builder /go/bin/todo_grpc_client /usr/local/bin/todo_grpc_client

CMD todo_grpc_client