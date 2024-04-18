package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sekolahbeta/pertemuan12/grpc/proto"
	"strings"

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

func (s *GreetService) GreetStream(req *proto.GreetStreamRequest,
	srv proto.GreetService_GreetStreamServer) error {

	log.Println("Memulai mengirim data")

	res := strings.Split(req.Username, ",")

	for i, usrname := range res {
		resp := proto.GreetStreamResponses{
			Part:    int64(i + 1),
			Message: fmt.Sprintf("Hallo %s", usrname),
		}
		if err := srv.Send(&resp); err != nil {
			log.Printf("Terjadi error mengirim data, err:%s\n",
				err.Error())
			return err
		}
	}
	return nil
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
