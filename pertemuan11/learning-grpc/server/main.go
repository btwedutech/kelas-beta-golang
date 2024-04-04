package main

import (
	"context"
	"fmt"
	"net"
	"sekolahbeta/pertemuan11/grpc/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GreetService struct {
	proto.GreetServiceServer
}

func (s *GreetService) Greet(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponses, error) {

	userName := req.Username
	phone := req.Phone

	response := fmt.Sprintf("Hello %s, no hp mu %d!", userName, phone)

	return &proto.GreetResponses{
		Greetresponse: response,
	}, nil

}

func main() {
	fmt.Println("Memulai gRPC Server ...")
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatalf("Gagal menginisialisasi server, error : %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	proto.RegisterGreetServiceServer(grpcServer, &GreetService{})
	grpcServer.Serve(grpcListener)
}
