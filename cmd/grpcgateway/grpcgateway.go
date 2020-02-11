package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	pb "github.com/ericatcorsha/rpc-compare/echo"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct{}

func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoReply, error) {
	// logrus.WithField("req", req).Info("Echo()")
	if req.Message == "user-error" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"The User Made an Error",
		)
	}
	return &pb.EchoReply{Message: req.Message}, nil
}

func runGRPC(wg sync.WaitGroup, lis net.Listener) {
	defer wg.Done()
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
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
	err := pb.RegisterEchoHandlerServer(ctx, gwmux, &server{})
	if err != nil {
		logrus.WithField("err", err).Fatal("GRPC Gateway could not register")
	}
	if err := http.Serve(lis, gwmux); err != nil {
		logrus.Fatal("GRPC Gateway could not serve")
	}

}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = ":8080"
	}
	l, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}

	logrus.WithField("listen", listen).Infof("Starting GRPC Service")
	go runGRPC(wg, l)

	listenGateway := os.Getenv("LISTEN_GATEWAY")
	if listenGateway == "" {
		listenGateway = ":8081"
	}
	lgw, err := net.Listen("tcp", listenGateway)
	if err != nil {
		log.Fatal(err)
	}
	logrus.WithField("listen", listenGateway).Infof("Starting GRPC Gateway Service")
	go runGateway(wg, lgw)

	wg.Wait()
}
