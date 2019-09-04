package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
	"./pb"
)

type HelloService struct {
}

func (hs HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello, %s", in.Username)}, nil
}

func (hs *HelloService) Chat(conn pb.HelloService_ChatServer)error {
	for {
		stream, err := conn.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println("receive from client:", stream.Stream)

		conn.Send(&pb.ServerStream {
		    Stream: newBytes(1,2,3,4,5),
	        })
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":6001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &HelloService{})
	go func() {
		s.Serve(lis)
	}()
	fmt.Println(s.GetServiceInfo())
	select {}
	//s.Serve(lis)
}

func newBytes(a ...byte)[]byte {
	return a
}

