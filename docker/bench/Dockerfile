FROM rpc-compare/base:latest

WORKDIR /src

# Install ghz for benchmarking the grpc service
RUN curl -L "https://github.com/bojand/ghz/releases/download/v0.49.0/ghz_0.49.0_Linux_x86_64.tar.gz" | tar xzv
RUN mv ./ghz /go/bin/
RUN mv ./ghz-web /go/bin/

# Install hey for benchmarking the REST services
#
RUN curl -L https://storage.googleapis.com/hey-release/hey_linux_amd64 > /go/bin/hey
RUN chmod +x /go/bin/hey

ADD ./docker/bench/bench.sh /src/
RUN chmod +x /src/bench.sh

# The volume for the output CSV files
VOLUME "/out"
CMD ["/src/bench.sh"]

