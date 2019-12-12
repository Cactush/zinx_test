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
	err := request.GetConnection().SendMsg(1, []byte("ping,ping,ping "))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("[zinx]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
