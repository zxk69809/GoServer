package service

type FunCode int32

//业务功能码定义
const (
	//更新设备状态
	FC_UP_STATUS FunCode = 1001
	//更新设备参数
	FC_UP_PARAM FunCode = 1002
	//心跳检测
	FC_HEARTBEAT FunCode = 1003
	//请求更新白名单
	FC_UP_WLIST FunCode = 1004
)
