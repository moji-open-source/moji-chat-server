package main

import (
	"log"
	"net"
	"strings"

	"github.com/moji-open-source/moji-chat-server/api/router"
	"github.com/moji-open-source/moji-chat-server/setup"
	"google.golang.org/grpc"
)

func main() {
	app := setup.App()
	defer app.Redis.Close()

	env := app.Env
	log.Println("env :", env)

	address := env.Address
	if strings.TrimSpace(address) == "" {
		address = ":8080"
	}

	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	router.Setup(router.RouterContext{
		Application: app,
		Server:      server,
	})

	log.Println("server run on", address, "address...")
	err = server.Serve(listen)

	if err != nil {
		panic(err)
	}
}
