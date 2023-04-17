package bot

import (
	"fmt"
	"strings"
)

type Server struct{
    UnimplementedBotServer
}

func (s *Server) SendMessage(stream Bot_SendMessageServer) error{
    for {
        in,err:= stream.Recv()
        if err != nil {
            return err
        }
        fmt.Printf("Received message  %s\n", in.Message)

        response := ChatMsg{}

        if strings.HasPrefix(in.Message,"/name"){
            response.Message = fmt.Sprintf("hello, %s", strings.TrimSpace(strings.TrimPrefix(in.Message, "/name")))
        } else{
            response.Message = "hello from the server!!"
        }

        if err:= stream.Send(&response); err!= nil{
            return err
        }
    }
}
