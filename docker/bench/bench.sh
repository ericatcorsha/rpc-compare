#!/usr/bin/env bash

echo >&2 "Waiting a sec to make sure the services are up"
sleep 1

# Benchmark the Baseline service
echo >&2 "Benchmarking Baseline Service"
hey -c 50 -n 10000 -d '{"message": "Hello, World!"}' -m POST -o csv http://baseline:8080/v1/echo > /out/baseline.csv
echo >&2 "Cooling down for 5 seconds"
sleep 5

# Benchmark the GRPC service
echo >&2 "Benchmarking GRPC Service"
ghz -c 50 -n 10000 --insecure --call=echo.Echo.Echo  -d '{"message": "Hello, World!"}' -O csv -o /out/grpc.csv grpcgateway:9090
echo >&2 "Cooling down for 5 seconds"
sleep 5

# Benchmark the GRPC-Gateway service
echo >&2 "Benchmarking GRPC-Gateway Service"
hey -c 50 -n 10000 -d '{"message": "Hello, World!"}' -m POST -o csv http://grpcgateway:9080/v1/echo > /out/grpc-gateway.csv
echo >&2 "Cooling down for 5 seconds"
sleep 5

# Benchmark the GRPC-Gateway service
echo >&2 "Benchmarking GraphQL Service"
hey -c 50 -n 10000 -d '{"operationName":"Echo","variables":{},"query":"mutation Echo { echo(input: {message: \"testing\"}) { message }}"}' -m POST -o csv http://graphql:7080/query > /out/graphql.csv
echo >&2 "Cooling down for 5 seconds"
sleep 5
