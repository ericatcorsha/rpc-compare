# A comparison of RPC methods

## Options

 - GRPC
 - GRPC-Gateway
 - GraphQL
 - HTTP + JSON + http.Handler

## Service

The service we're implementing is an echo service.  If the incoming message is
the string "user-error", the standard validation error for the solution will be returned

## Benchmarking

Open in a terminal:

``` sh
make up
```

Open in another terminal

``` sh
make bench
```
