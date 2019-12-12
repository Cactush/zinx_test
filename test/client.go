package main

import (
	"fmt"
	"github.com/Cactush/zinx_test/znet"
	"io"
	"net"
	"time"
)

/*
   模拟客户端
*/
func main() {

	fmt.Println("Client Test ... start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	for {
		dp := znet.NewDataPack()
		msgpck := znet.NewMsgPackage(0, []byte("zinx test message"))
		fmt.Println(msgpck.Id, msgpck.DataLen)
		msg, _ := dp.Pack(msgpck)
		fmt.Println(msg)
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}
		headData := make([]byte, dp.GetHeadLen())

		_, err = io.ReadFull(conn, headData)
		fmt.Println(headData)
		if err != nil {
			fmt.Println("read head error")
			break
		}
		print(headData)
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err: ", err)
			return
		}
		if msgHead.GetDatalen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDatalen())
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err: ", err)
				return
			}
			fmt.Println("==> Recv Msg: ID= ", msg.Id, "data= ", string(msg.Data))
		}
		time.Sleep(1 * time.Second)
	}

}
