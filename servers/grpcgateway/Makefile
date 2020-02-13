RUN=docker run -it -v `pwd`:/src rpc-compare/base

protoc:
	mkdir -p echo
	${RUN} protoc ./echo.proto \
	  -I . \
	  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  --go_out=plugins=grpc:echo/. \
	  --grpc-gateway_out=logtostderr=true:echo/. \
	  --swagger_out=logtostderr=true:.
