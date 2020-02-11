FROM rpc-compare/base:latest

WORKDIR /src

ADD go.mod /src/
ADD go.sum /src/
RUN go mod download

ADD gqlgen.yml /src/
ADD generated.go /src/
ADD models_gen.go /src/
ADD resolver.go /src/
ADD schema.graphql /src/
ADD server /src/server

RUN go test ./...
RUN go build -o server ./server/

CMD ["./server"]
