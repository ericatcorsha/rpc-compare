FROM rpc-compare/base:latest

WORKDIR /src

ADD servers/grpcgateway/ /src/
RUN go build -o grpcgateway ./

CMD ["./grpcgateway"]
