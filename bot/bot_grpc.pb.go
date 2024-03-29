// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bot/bot.proto

package bot

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Bot_SendMessage_FullMethodName = "/bot.Bot/SendMessage"
)

// BotClient is the client API for Bot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotClient interface {
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (Bot_SendMessageClient, error)
}

type botClient struct {
	cc grpc.ClientConnInterface
}

func NewBotClient(cc grpc.ClientConnInterface) BotClient {
	return &botClient{cc}
}

func (c *botClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Bot_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bot_ServiceDesc.Streams[0], Bot_SendMessage_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &botSendMessageClient{stream}
	return x, nil
}

type Bot_SendMessageClient interface {
	Send(*ChatMsg) error
	Recv() (*ChatMsg, error)
	grpc.ClientStream
}

type botSendMessageClient struct {
	grpc.ClientStream
}

func (x *botSendMessageClient) Send(m *ChatMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *botSendMessageClient) Recv() (*ChatMsg, error) {
	m := new(ChatMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BotServer is the server API for Bot service.
// All implementations must embed UnimplementedBotServer
// for forward compatibility
type BotServer interface {
	SendMessage(Bot_SendMessageServer) error
	mustEmbedUnimplementedBotServer()
}

// UnimplementedBotServer must be embedded to have forward compatible implementations.
type UnimplementedBotServer struct {
}

func (UnimplementedBotServer) SendMessage(Bot_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedBotServer) mustEmbedUnimplementedBotServer() {}

// UnsafeBotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotServer will
// result in compilation errors.
type UnsafeBotServer interface {
	mustEmbedUnimplementedBotServer()
}

func RegisterBotServer(s grpc.ServiceRegistrar, srv BotServer) {
	s.RegisterService(&Bot_ServiceDesc, srv)
}

func _Bot_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BotServer).SendMessage(&botSendMessageServer{stream})
}

type Bot_SendMessageServer interface {
	Send(*ChatMsg) error
	Recv() (*ChatMsg, error)
	grpc.ServerStream
}

type botSendMessageServer struct {
	grpc.ServerStream
}

func (x *botSendMessageServer) Send(m *ChatMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *botSendMessageServer) Recv() (*ChatMsg, error) {
	m := new(ChatMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Bot_ServiceDesc is the grpc.ServiceDesc for Bot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bot.Bot",
	HandlerType: (*BotServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _Bot_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bot/bot.proto",
}
