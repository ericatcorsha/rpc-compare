FROM rpc-compare/base:latest

WORKDIR /src

ADD servers/baseline/ /src/
RUN go build -o baseline ./

CMD ["./baseline"]
