package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/moji-open-source/moji-chat-server/hello"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	pb.UnimplementedSayHelloServer
}

func (*GrpcServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	log.Printf("received grpc req: %+v", req.String())
	return &pb.HelloResp{Msg: fmt.Sprintf("hello world! %s", req.Name)}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	pb.RegisterSayHelloServer(server, &GrpcServer{})

	log.Println("server run on 8081 port...")
	err = server.Serve(listen)

	if err != nil {
		panic(err)
	}
}
