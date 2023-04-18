package bot

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestServer_SendMessage(t *testing.T) {
    // Set up a test server
    s := grpc.NewServer()
    defer s.Stop()

    // Register your gRPC service implementation with the test server
    chatService := &Server{}
    RegisterBotServer(s, chatService)

    // Start the test server in a separate goroutine
    listener, err := net.Listen("tcp", ":9000")
    if err != nil {
        t.Fatalf("Failed to start test server: %v", err)
    }
    go s.Serve(listener)
    defer listener.Close()

    // Create a client connection to the test server
    conn, err := grpc.Dial(listener.Addr().String(), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to connect to test server: %v", err)
    }
    defer conn.Close()

    // Create a client instance using the connection
    client := NewBotClient(conn)

    // Call the SendMessage method with a test message

    // Verify that the response message matches the expected value

    stream,err := client.SendMessage(context.Background())
    require.NoError(t,err)

    message1 := &ChatMsg{
        Message: "Hello, world!",
    }

    err = stream.Send(message1)

    require.NoError(t,err)

    message2 := &ChatMsg{Message: "/name Gaurav"}

    err = stream.Send(message2)

    require.NoError(t,err)

    stream.CloseSend()
    response,err := stream.Recv()
    require.NoError(t,err)
    assert.Equal(t,"hello from the server!!",response.Message)

    response,err = stream.Recv()
    require.NoError(t,err)
    assert.Equal(t,"hello, Gaurav",response.Message)

}
