package service

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"goserver/src/app/model"
	"goserver/src/library/protocol"
)

type BllUpStatus struct {
	ver string
}

func init() {
	RegBllMode(FC_UP_STATUS, BllUpStatus{})
}

func (b *BllUpStatus) ThisName() string {
	return "更新状态"
}

func (b *BllUpStatus) RunBll(pack protocol.MsgPack) (protocol.MsgPack, error) {
	var err error
    r := model.GetAllUser()
	if r != nil {
		u := model.User{}
		for _,v := range r {
			if err := v.ToStruct(&u); err == nil {
				fmt.Println(" uid:", u.Id, "name:", u.Nickname)
			} else {
				fmt.Println(err)
			}
		}
	}
	glog.Debugf("BllUpStatus RunBll %s", gtime.Datetime())
	return pack, err
}
