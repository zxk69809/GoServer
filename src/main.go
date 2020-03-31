package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	_ "goserver/src/boot"
)

func main() {
	err := g.TCPServer().Run()
	if err != nil {
		glog.Errorf("服务启动失败%s", err)
	}
}
