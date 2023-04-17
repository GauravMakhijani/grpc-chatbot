package main

import (
	"fmt"
	"grpc_chatbot/bot"
	"net"

	"google.golang.org/grpc"
)

func main() {
    lis,err:= net.Listen("tcp",":9000")
    if err != nil {
        fmt.Printf("failed to listen : %v",err)
        return
    }

    s:= grpc.NewServer()

    bot.RegisterBotServer(s,&bot.Server{})

    fmt.Println("starting server at :9000")

    if err:= s.Serve(lis); err != nil {
        fmt.Printf("failed to serve : %v",err)
        return
    }
}
