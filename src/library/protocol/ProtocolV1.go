package protocol

import (
	"encoding/hex"
	"github.com/gogf/gf/util/gconv"
)

type ProtocolV1 struct {
	ver string
}

const (
	TCP_HEADER     = "~"
	TCP_HEADER_LEN = 1
	TCP_DATA_LEN   = 2
)

// 封包
func (p *ProtocolV1) Packet(msg MsgPack) []byte {
	return append(append([]byte(TCP_HEADER), IntToBytes(msg.Len)...), msg.Data...)
}

// 解包
func (p *ProtocolV1) Unpack(buffer []byte, readerChannel chan MsgPack) []byte {
	length := len(buffer)
	var i int
	for i = 0; i < length; i++ {
		if length < i+TCP_HEADER_LEN+TCP_DATA_LEN {
			break
		}
		if string(buffer[i:i+TCP_HEADER_LEN]) == TCP_HEADER {
			msgLen := BytesToInt(buffer[i+TCP_HEADER_LEN : i+TCP_HEADER_LEN+TCP_DATA_LEN])
			if length < i+TCP_HEADER_LEN+TCP_DATA_LEN+msgLen {
				break
			}
			data := buffer[i+TCP_HEADER_LEN+TCP_DATA_LEN : i+TCP_HEADER_LEN+TCP_DATA_LEN+msgLen]
			fun := gconv.Int(hex.EncodeToString(data[0:2]))
			load := data[2:]
			msg := MsgPack{
				Len:  len(load),
				Fun:  fun,
				Data: load,
				Mac:  "abc",
			}
			readerChannel <- msg

			i += TCP_HEADER_LEN + TCP_DATA_LEN + msgLen - 1
		}
	}
	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}
