package service

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"goserver/src/library/protocol"
)

type BllUpParam struct {
	ver string
}

func init() {
	RegBllMode(FC_UP_PARAM, BllUpParam{})
}

func (b *BllUpParam) ThisName() string {
	return "更新参数"
}

func (b *BllUpParam) RunBll(pack protocol.MsgPack) (protocol.MsgPack, error) {
	var err error
	glog.Debugf("BllUpParam RunBll %s", gtime.Datetime())
	return pack, err
}
