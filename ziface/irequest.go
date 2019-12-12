package ziface

/*
	IRequest接口
	实际上是吧客户端请求的链接信息和请求信息放到Request中
*/

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
}
