package test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"testing"

	pb "github.com/ericatcorsha/rpc-compare/echo"
	"google.golang.org/grpc"
)

func BenchmarkGRPC(b *testing.B) {
	// Create client
	conn, err := grpc.Dial(os.Getenv("GRPC_SERVER"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewEchoClient(conn)

	for i := 0; i < b.N; i++ {
		// TODO: Make call
		_, err := client.Echo(context.Background(), &pb.EchoRequest{Message: "Hello, World!"})
		if err != nil {
			b.Error(err)
		}
	}
}

type EchoReq struct {
	Message string `json:"message"`
}

type EchoReply struct {
	Message string `json:"message"`
}

func CallRestEcho(endpoint string) (*EchoReply, error) {
	var reqBuf strings.Builder
	req := EchoReq{Message: "Hello, World!"}
	if err := json.NewEncoder(&reqBuf).Encode(req); err != nil {
		return nil, err
	}
	res, err := http.Post(endpoint, "application/json", strings.NewReader(reqBuf.String()))
	if err != nil {
		return nil, err
	}

	reply := EchoReply{}
	if err := json.NewDecoder(res.Body).Decode(&reply); err != nil {
		return nil, err
	}
	return &reply, nil
}

func BenchmarkGateway(b *testing.B) {
	endpoint := os.Getenv("GRPCGATEWAY_ECHO_URL")
	for i := 0; i < b.N; i++ {
		_, err := CallRestEcho(endpoint)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkBaseline(b *testing.B) {
	endpoint := os.Getenv("BASELINE_ECHO_URL")
	for i := 0; i < b.N; i++ {
		_, err := CallRestEcho(endpoint)
		if err != nil {
			b.Error(err)
		}
	}
}
