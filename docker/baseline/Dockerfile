FROM rpc-compare/base:latest

WORKDIR /src

ADD go.mod /src/
ADD go.sum /src/
RUN go mod download

ADD cmd /src/cmd

RUN go build -o server ./cmd/baseline

CMD ["./server"]
