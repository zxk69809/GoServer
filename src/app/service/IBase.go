package service

import "goserver/src/library/protocol"

type IBase interface {
	//功能名称
	ThisName() string
	//业务处理
	RunBll(pack protocol.MsgPack) (protocol.MsgPack, error)
}
