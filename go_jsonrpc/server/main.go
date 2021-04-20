package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(request string,reply *string) error{
	*reply = "hello new world! req = " + request
	fmt.Printf("Server get a request.\n")
	return nil
}

func main(){
	rpc.RegisterName("HelloService",new(HelloService))
	listen,err := net.Listen("tcp",":12345")
	if err != nil{
		log.Fatalf("Listen tcp err: %v\n",err)
	}
	conn,err := listen.Accept()
	if err != nil{
		log.Fatalf("Accept err: %v",err)
	}
	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

}
