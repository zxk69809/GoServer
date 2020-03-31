package protocol

import (
	"bytes"
	"encoding/binary"
)

type MsgPack struct {
	//长度
	Len int
	//类型
	Fun int
	//数据
	Data []byte
	//校验码
	Mac string
}

type IParser interface {
	//协议组包
	Packet(msg MsgPack) []byte
	//协议解包
	Unpack(buffer []byte, readerChannel chan MsgPack) []byte
}

// 整形转换成字节
func IntToBytes(n int) []byte {
	x := int16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int16
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}
