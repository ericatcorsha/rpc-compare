RUN=docker run -it -v `pwd`:/src rpc-compare/base

all: clean protoc grpcgateway baseline

protoc: docker-base
	# grpcgateway
	$(MAKE) -C servers/grpcgateway protoc
	$(MAKE) -C servers/twirp protoc


docker-base:
	docker-compose build base

grpcgateway:
	docker-compose build grpcgateway

baseline:
	docker-compose build baseline

graphql:
	docker-compose build graphql

go-swagger: docker-base
	$(MAKE) -C servers/go-swagger
	docker-compose build go-swagger

up: protoc
	docker-compose up --build --remove-orphans

sizes:
	echo "baseline:"
	docker run rpc-compare/baseline ls -lh /src/server
	echo "grpcgateway:"
	docker run rpc-compare/grpcgateway ls -lh /src/server
