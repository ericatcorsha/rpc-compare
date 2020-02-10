FROM todo-grpc/base

WORKDIR /src

ADD go.mod /src/
ADD go.sum /src/
RUN go mod download

ADD main.go /src/
ADD todo /src/todo

RUN go test ./...
RUN go build -o server .

ADD todo.swagger.json ./
CMD ["./server"]
