package main

import (
	"fmt"
	"github.com/Cactush/zinx_test/ziface"
	"github.com/Cactush/zinx_test/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router Prehandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ..."))
	if err != nil {
		fmt.Println("Call back ping ping ping error")
	}
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping..."))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	s := znet.NewServer("[zinx]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
