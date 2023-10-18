package main

import (
	"context"
	"fmt"
	"log"

	"com.wlq/simplebank/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	username = "test1"
	password = "123456"
	bearer   = "Bearer "
)

var accessToken string = ""

func main() {
	serverAddress := "localhost:7778"

	cc, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	server := pb.NewSimpleBankClient(cc)
	TestLoginUser(server)
	TestUpdateUser(server)

}

func TestLoginUser(server pb.SimpleBankClient) {
	ctx := context.Background()
	rsp, err := server.LoginUser(ctx, &pb.LoginUserRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}
	accessToken = rsp.GetAccessToken()
	fmt.Println(protojson.Format(rsp))
}

func TestUpdateUser(server pb.SimpleBankClient) {
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", bearer+accessToken)
	req := &pb.UpdateUserRequest{
		Username: username,
	}
	rsp, err := server.UpdateUser(ctx, req)
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}
	fmt.Println(protojson.Format(rsp))
}
