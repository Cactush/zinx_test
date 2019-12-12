package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/Cactush/zinx_test/utils"
	"github.com/Cactush/zinx_test/ziface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	return 8
}

// 封包
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	databuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(databuff, binary.LittleEndian, msg.GetDatalen()); err != nil {
		return nil, err
	}
	if err := binary.Write(databuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(databuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return databuff.Bytes(), nil
}

//拆包
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	fmt.Println(binaryData)
	dataBuff := bytes.NewReader(binaryData)
	msg := &Message{}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("Too large msg data received")
	}
	return msg, nil
}
