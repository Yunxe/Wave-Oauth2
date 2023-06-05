package rpc

import (
	pb "oauth2/proto/sso_client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Grpc pb.TokenServiceClient

func InitGrpcClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("grpc dial err:" + err.Error())
	}

	//defer conn.Close()

	Grpc = pb.NewTokenServiceClient(conn)
}
