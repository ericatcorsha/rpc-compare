package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/ericatcorsha/rpc-compare/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

const (
	endpoint = ":8080"
)

type server struct{}

func runGRPC(wg sync.WaitGroup, lis net.Listener) {
	defer wg.Done()
	s := grpc.NewServer()
	service.RegisterEchoServer(s, &server{})
	logrus.WithField("port", endpoint).Infof("Starting GRPC Service")
	if err := s.Serve(lis); err != nil {
		logrus.Fatal("GRPC could not serve")
	}
}

func runGateway(wg sync.WaitGroup, lis net.Listener) {
	defer wg.Done()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	err := service.RegisterEchoHandlerServer(ctx, gwmux, &server{})
	if err != nil {
		logrus.WithField("err", err).Fatal("GRPC Gateway could not register")
	}
	logrus.WithField("port", endpoint).Infof("Starting GRPC Gateway Service")
	if err := http.Serve(lis, gwmux); err != nil {
		logrus.Fatal("GRPC Gateway could not serve")
	}

}

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(2)

	l, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(l)

	go runGRPC(wg, m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc")))
	go runGateway(wg, m.Match(cmux.HTTP1Fast()))
	wg.Wait()
}
