package utils

import (
	"encoding/json"
	"github.com/Cactush/zinx_test/ziface"
	"io/ioutil"
)

type GlobalObj struct {
	TcpServer     ziface.IServer
	Host          string
	TcpPort       int
	Name          string
	Version       string
	MaxPacketSize uint32
	MaxConn       int
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
func init() {
	GlobalObject = &GlobalObj{
		Host:          "0.0.0.0",
		TcpPort:       7777,
		Name:          "ZinxServer",
		Version:       "v0.4",
		MaxPacketSize: 4096,
		MaxConn:       12000,
	}

	GlobalObject.Reload()
}