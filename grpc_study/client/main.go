package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"./pb"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:6001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Username: "ft"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r.Message)

	chatClient, e := c.Chat(context.Background())
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	go func(){
		for {
			stream, e:=chatClient.Recv()
	                if e == io.EOF {
		          fmt.Println("EOF")
		          return
	                }
	                if e != nil {
		          //fmt.Println(errorx.Wrap(e).Error())
		          fmt.Println("There is errorx")
		          return
	                }
	                fmt.Println("receive from server:", stream.Stream)
		}
	}()
	chatClient.Send(&pb.ClientStream{
	  Stream: newBytes(10,9,8,7),
        })
	select{
	case <-time.After(20*time.Second):
		fmt.Println("Test ended!!!")
	}
}

func newBytes(a ...byte)[]byte{
	return a
}

