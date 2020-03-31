package api

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"goserver/src/app/service"
	"goserver/src/library/protocol"
)

//协议处理
var sipp protocol.IParser

//客户端连接池
var ConnMapping = gmap.NewAnyAnyMap()

func init() {
	ver := g.Config().GetInt("server.protover")
	switch ver {
	case 1:
		sipp = new(protocol.ProtocolV1)
	default:
		sipp = new(protocol.ProtocolV1)
	}
}

//接收处理
func reader(addr string, readerChannel chan protocol.MsgPack) {
	for {
		select {
		case msg, ok := <-readerChannel:
			if !ok {
				glog.Debugf("接收处理通道关闭...")
				return
			}
			glog.Debugf(string(msg.Data))
			if ConnMapping.Contains(addr) {
				connHand, ok := ConnMapping.Get(addr).(*TcpConn)
				if ok {
					//根据功能码实例化对应业务模块
					bll, err := service.NewBllMode(service.FunCode(msg.Fun))
					if err != nil {
						glog.Errorf("NewBllMode Error：%s", err.Error())
					} else {
						//执行具体业务
						resp, err := bll.(service.IBase).RunBll(msg)
						if err != nil {
							glog.Errorf("RunBll Error：%s", err.Error())
						} else {
							tmpBuffer := sipp.Packet(resp)
							if err := connHand.Conn().Send(tmpBuffer); err != nil {
								glog.Errorf("conn.Send Error：%s", err.Error())
							}
						}
					}
				}
			}
		}
	}
}

//tcp监听处理
func Tcphandler(conn *gtcp.Conn) {
	defer conn.Close()
	glog.Debugf("接受一个客户端连接,%s", conn.RemoteAddr())
	ConnMapping.Set(conn.RemoteAddr().String(), NewTcpConn("", conn))
	glog.Debugf("当前客户端连接池：%d", ConnMapping.Size())
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)
	//声明一个管道用于接收解包的数据
	readerChannel := make(chan protocol.MsgPack, 16)
	//结束关闭接收管道
	defer close(readerChannel)
	go reader(conn.RemoteAddr().String(), readerChannel)
	for {
		data, err := conn.Recv(-1)
		if len(data) > 0 {
			tmpBuffer = sipp.Unpack(append(tmpBuffer, data...), readerChannel)
		}
		if err != nil {
			break
		}
	}
	//结束清理客户端连接池
	ConnMapping.Remove(conn.RemoteAddr().String())
	glog.Debugf("当前客户端连接池：%d", ConnMapping.Size())
	glog.Debugf("断开一个客户端连接,%s", conn.RemoteAddr())
}
