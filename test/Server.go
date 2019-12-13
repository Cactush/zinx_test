package main

import (
	"fmt"
	"github.com/Cactush/zinx_test/ziface"
	"github.com/Cactush/zinx_test/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	fmt.Println("recv from client:msgId =", request.GetMsgID(), "data= ", request.GetData())
	err := request.GetConnection().SendMsg(0, []byte("ping,ping,ping "))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	fmt.Println("recv from client: msgId= ", request.GetMsgID(), ", data= ", request.GetData())
	err := request.GetConnection().SendMsg(1, []byte("hello zinx router"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("[zinx]")
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Serve()
}
