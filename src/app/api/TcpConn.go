package api

import (
	"github.com/gogf/gf/net/gtcp"
	"github.com/rs/xid"
)

//客户端Tcp连接
type TcpConn struct {
	guid   xid.ID      //唯一标识
	termSN string      //终端编号
	param  interface{} //任意关联参数
	conn   *gtcp.Conn  //终端tcp连接
}

func (p *TcpConn) Conn() *gtcp.Conn {
	return p.conn
}
func (p *TcpConn) TermSN() string {
	return p.termSN
}

func (p *TcpConn) Guid() xid.ID {
	return p.guid
}

func (p *TcpConn) SetTermSN(termSN string) {
	p.termSN = termSN
}

func NewTcpConn(sn string, con *gtcp.Conn) *TcpConn {
	return &TcpConn{
		guid:   xid.New(),
		termSN: sn,
		conn:   con,
	}
}

func (p *TcpConn) SetParam(param interface{}) {
	p.param = param
}

func (p *TcpConn) GetParam() interface{} {
	return p.param
}
