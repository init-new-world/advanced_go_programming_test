package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main(){
	conn,err := net.Dial("tcp","localhost:12345")
	if err != nil{
		log.Fatalf("Dial tcp err: %v\n",err)
	}

	client:= rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello","initialize",&reply)
	if err != nil{
		log.Fatalf("Client err: %v\n",err)
	}
	fmt.Println(reply)
}
