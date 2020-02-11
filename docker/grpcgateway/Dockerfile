FROM rpc-compare/base:latest

WORKDIR /src

ADD go.mod /src/
ADD go.sum /src/
RUN go mod download

ADD cmd /src/cmd
ADD echo /src/echo

RUN go build -o server ./cmd/grpcgateway

ADD service.swagger.json ./
CMD ["./server"]
