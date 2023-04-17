package main

import (
	"context"
	"fmt"
	"grpc_chatbot/bot"

	"google.golang.org/grpc"
)

func main(){
    conn,err:= grpc.Dial(":9000",grpc.WithInsecure())
    if err != nil {
        fmt.Printf("failed to connect: %v",err)
        return
    }

    defer conn.Close()

    c:= bot.NewBotClient(conn)

    stream,err := c.SendMessage(context.Background())
    if err != nil {
        fmt.Printf("Error sending message: %v",err)
        return
    }

    message1 := bot.ChatMsg{Message: "Hello from client!!"}

    if err:= stream.Send(&message1); err!=nil{
        fmt.Printf("Error sending message1: %v",err)
        return
    }

    message2 := bot.ChatMsg{Message: "/name Gaurav"}

    if err:= stream.Send(&message2); err!=nil{
        fmt.Printf("Error sending message2: %v",err)
        return
    }
    stream.CloseSend()
    waitc := make(chan struct{})
    go func ()  {
        for{
            response,err:= stream.Recv()
            if err != nil {
                fmt.Printf("Error receiving message: %v", err)
                close(waitc)
                return
            }
            fmt.Printf("recieved msg from server : %s\n",response.Message)
        }
    }()
    <-waitc

}
