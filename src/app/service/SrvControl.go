package service

import (
	"errors"
	"github.com/gogf/gf/container/gmap"
	"reflect"
)

//业务管理
type BllControl struct {
	//终端编号
	TermID int32
	//业务模块
	bllMode IBase
}

//业务关系集合
var typeRegistryMap = gmap.NewAnyAnyMap()

//注册业务对象
func RegBllMode(fun FunCode, elem interface{}) {
	t := reflect.ValueOf(elem).Type()
	typeRegistryMap.Set(fun, t)
}

//实例化业务对象
func NewBllMode(fun FunCode) (interface{}, error) {
	if typeRegistryMap.Contains(fun) {
		t := typeRegistryMap.Get(fun).(reflect.Type)
		return reflect.New(t).Interface(), nil
	}
	return nil, errors.New("未找到对应的业务功能")
}
