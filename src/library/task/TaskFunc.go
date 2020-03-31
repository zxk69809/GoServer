package task

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"goserver/src/app/api"
)

func TaskUpStatus() {
	glog.Println("TaskUpStatus run ")
}

func TaskHeartbeat() {
	api.ConnMapping.Iterator(func(k interface{}, v interface{}) bool {
		tcpConn := v.(*api.TcpConn)
		send := fmt.Sprintf("hello %d", tcpConn.Guid().Counter())
		tcpConn.Conn().Send([]byte(send))
		return true
	})
}
