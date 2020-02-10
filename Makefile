RUN=docker run -it -v `pwd`:/src todo-grpc/base

all: clean protoc grpcgateway graphql baseline

clean:
	rm -f grpcgateway graphql baseline

protoc:
	mkdir -p echo
	${RUN} protoc echo.proto \
	  -I . \
	  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  --go_out=plugins=grpc:echo \
	  --grpc-gateway_out=logtostderr=true:echo/. \
	  --swagger_out=logtostderr=true:.

grpcgateway:
	go build -o grpcgateway grpcgateway.go

graphql:
	go build -o graphql graphql.go

baseline:
	go build -o baseline baseline.go

bench:
	false
