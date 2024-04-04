package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sekolahbeta/pertemuan11/grpc/proto"
	"strings"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		logrus.Fatalf("Gagal menginisialisasi client, error : %s", err.Error())
	}
	defer conn.Close()

	greetService := proto.NewGreetServiceClient(conn)

	fmt.Println("Masukkan Username :")

	usrNameInput := bufio.NewReader(os.Stdin)
	usrName, err := usrNameInput.ReadString('\n')
	if err != nil {
		logrus.Fatalf("Terjadi Error: %s", err.Error())
		return
	}

	usrName = strings.Replace(usrName, "\n", "", 1)
	usrName = strings.Replace(usrName, "\r", "", 1)

	var phoneNumber int64

	fmt.Println("Masukkan No HP :")
	_, err = fmt.Scanln(&phoneNumber)
	if err != nil {
		logrus.Fatalf("Terjadi Error: %s", err.Error())
		return
	}
	greetResponse, err := Greet(context.Background(),
		greetService, usrName,
		phoneNumber)
	if err != nil {
		logrus.Fatalf("Terjadi error saat melakukan greetings, err:%s",
			err.Error())
	}
	fmt.Println(" ====================== ")
	fmt.Println(greetResponse)
}

func Greet(ctx context.Context,
	client proto.GreetServiceClient, name string, phoneNumber int64) (string, error) {
	res, err := client.Greet(ctx, &proto.GreetRequest{
		Username: name,
		Phone:    phoneNumber,
	})
	if err != nil {
		return "", err
	}
	return res.Greetresponse, nil
}
