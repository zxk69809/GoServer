package service

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"goserver/src/library/protocol"
)

type BllHeartbeat struct {
	ver string
}

func init() {
	RegBllMode(FC_HEARTBEAT, BllHeartbeat{})
}

func (b *BllHeartbeat) ThisName() string {
	return "心跳检查"
}

func (b *BllHeartbeat) RunBll(pack protocol.MsgPack) (protocol.MsgPack, error) {
	var err error
	glog.Debugf("BllHeartbeat RunBll %s", gtime.Datetime())
	return pack, err
}
