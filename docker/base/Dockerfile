FROM golang

WORKDIR /src

RUN apt update
RUN apt install -y protobuf-compiler

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go


RUN curl -L https://github.com/grpc/grpc-web/releases/download/1.0.7/protoc-gen-grpc-web-1.0.7-linux-x86_64 > /usr/bin/protoc-gen-grpc-web
RUN chmod +x /usr/bin/protoc-gen-grpc-web

RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go get -u github.com/a-urth/go-bindata/...

# Tools for the grpc-gateway swagger client
RUN apt install -y nodejs npm
RUN npm install swagger-typescript-client-generator

# Install ghz for benchmarking the grpc service
WORKDIR /tmp
RUN curl -L "https://github.com/bojand/ghz/releases/download/v0.49.0/ghz_0.49.0_Linux_x86_64.tar.gz" | tar xzv
RUN mv ./ghz /go/bin/
RUN mv ./ghz-web /go/bin/

WORKDIR /src
