RUN=docker run -it -v `pwd`:/src rpc-compare/base

all: clean protoc swagger-gen grpcgateway baseline

protoc: docker-base
	$(MAKE) -C servers/grpcgateway protoc


docker-base:
	docker-compose build base

grpcgateway:
	docker-compose build grpcgateway

baseline:
	docker-compose build baseline

graphql:
	docker-compose build graphql

swagger-gen: docker-base
	$(MAKE) -C servers/go-swagger

build:
	docker-compose build

up: protoc build benchmark-binary-size
	docker-compose up --build --remove-orphans

benchmark-binary-size:
	# record the sizes of the docker images
	echo >&2 "Recording the size of the compiled binaries"
	echo "binary,size (bytes)" > out/bin-sizes.csv
	docker run rpc-compare/baseline stat -c "%n,%s" baseline >> out/bin-sizes.csv
	docker run rpc-compare/go-swagger stat -c "%n,%s" go-swagger >> out/bin-sizes.csv
	docker run rpc-compare/graphql stat -c "%n,%s" graphql >> out/bin-sizes.csv
	docker run rpc-compare/grpcgateway stat -c "%n,%s" grpcgateway >> out/bin-sizes.csv
