package main

import (
	"log"
	"net"

	"github.com/moji-open-source/moji-chat-server/abi/user"
	"github.com/moji-open-source/moji-chat-server/db"
	server_user "github.com/moji-open-source/moji-chat-server/server/user"
	"github.com/moji-open-source/moji-chat-server/setup"
	"google.golang.org/grpc"
)

type GrpcUserServer struct {
	user.UnimplementedUserServer
}

func main() {
	app := setup.App()

	env := app.Env

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	server_user.RegisterServer(server)

	log.Println("server run on 8081 port...")
	err = server.Serve(listen)

	if err != nil {
		panic(err)
	}
}

func beforeSetup() {
	db.InitDatabase()
}
