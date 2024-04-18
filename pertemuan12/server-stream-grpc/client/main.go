package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sekolahbeta/pertemuan12/grpc/proto"
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

	usrNames := []string{}

	for i := 0; i < 5; i++ {
		fmt.Println("Masukkan Username :")
		usrNameInput := bufio.NewReader(os.Stdin)
		usrName, err := usrNameInput.ReadString('\n')
		if err != nil {
			logrus.Fatalf("Terjadi Error: %s", err.Error())
			return
		}
		usrName = strings.Replace(usrName, "\n", "", 1)
		usrName = strings.Replace(usrName, "\r", "", 1)

		usrNames = append(usrNames, usrName)
	}

	err = GreetStream(context.Background(),
		greetService, usrNames)
	if err != nil {
		logrus.Printf("Terjadi error request streaming, error : %s",
			err.Error())
	}
	// var phoneNumber int64

	// fmt.Println("Masukkan No HP :")
	// _, err = fmt.Scanln(&phoneNumber)
	// if err != nil {
	// 	logrus.Fatalf("Terjadi Error: %s", err.Error())
	// 	return
	// }
	// greetResponse, err := Greet(context.Background(),
	// 	greetService, usrName,
	// 	phoneNumber)
	// if err != nil {
	// 	logrus.Fatalf("Terjadi error saat melakukan greetings, err:%s",
	// 		err.Error())
	// }
	// fmt.Println(" ====================== ")
	// fmt.Println(greetResponse)
}

func GreetStream(ctx context.Context,
	client proto.GreetServiceClient, usrNames []string) error {
	usrName := ""

	for i, name := range usrNames {
		if i > 0 {
			usrName += fmt.Sprintf(",%s", name)
		} else {
			usrName = name
		}
	}
	stream, err := client.GreetStream(ctx,
		&proto.GreetStreamRequest{
			Username: usrName,
		})

	if err != nil {
		logrus.Fatalf("Terjadi error saat melakukan stream greetings, err:%s",
			err.Error())
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err == nil {
			fmt.Printf("Respon %d : %s\n", resp.Part, resp.Message)
		} else {
			logrus.Printf("Terjadi error saat streaming, err :%s\n",
				err.Error())
			return err
		}
	}

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
