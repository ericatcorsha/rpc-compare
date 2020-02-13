FROM rpc-compare/base:latest

WORKDIR /src

ADD servers/go-swagger/ /src/
RUN go build -o go-swagger ./cmd/echo-server

CMD ["./go-swagger"]
