package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtimer"
	api_tcpserver "goserver/src/app/api"
	"goserver/src/library/task"
	"time"
)

// 用于应用初始化。
func init() {
	//开启debug模式
	g.SetDebug(true)
	c := g.Config()
	glog.Infof("Ver 1.0.0 20190505")
	glog.Infof("初始化配置信息...")
	s := g.TCPServer()
	glog.Infof("初始化服务对象...")
	// glog配置
	logpath := c.GetString("setting.logpath")
	glog.Infof("日志文件路径：%s", logpath)
	glog.SetPath(logpath)
	glog.SetStdoutPrint(true)

	// Tcp Server配置
	address := c.GetString("server.address")
	glog.Debugf("服务绑定地址和端口：%s", address)
	s.SetAddress(address)
	s.SetHandler(api_tcpserver.Tcphandler)
	glog.Debugf("服务启动完毕")

	//启动定时任务
	//指定时间点的任务
	//	gcron.AddSingleton("@every 1m", task.TaskUpStatus)
	//	g.Dump(gtimer.Entries())
	//指定间隔时间任务
	interval := time.Second * 5
	gtimer.AddSingleton(interval, task.TaskHeartbeat)


}
