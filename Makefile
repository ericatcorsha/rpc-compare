RUN=docker run -it -v `pwd`:/src rpc-compare/base

all: clean protoc grpcgateway baseline

protoc: docker-base
	mkdir -p echo
	${RUN} protoc echo.proto \
	  -I . \
	  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  --go_out=plugins=grpc:echo \
	  --grpc-gateway_out=logtostderr=true:echo/. \
	  --swagger_out=logtostderr=true:.

docker-base:
	docker-compose build base

grpcgateway:
	docker-compose build grpcgateway

baseline:
	docker-compose build baseline

up: protoc
	docker-compose up --build --remove-orphans

sizes:
	echo "baseline:"
	docker run rpc-compare/baseline ls -lh /src/server
	echo "grpcgateway:"
	docker run rpc-compare/grpcgateway ls -lh /src/server
