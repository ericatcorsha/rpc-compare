FROM rpc-compare/base:latest

WORKDIR /src

ADD servers/graphql/ /src/
RUN go build -o graphql ./server

CMD ["./graphql"]
