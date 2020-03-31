package model

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

type User struct {
	Id  int
	Passport string
	Password string
	Nickname string
	Createtime string
}
var (
	// 表对象
	table = g.DB().Table("user").Safe()
)

func GetUser(id int) gdb.Record{
	if r, err := table.Where("id = ?", id).One(); err == nil {
		return  r
	} else {
		fmt.Println(err)
	}
	return  nil
}

func GetAllUser() gdb.Result{
	if r, err := table.Select(); err == nil {
		return  r
	} else {
		fmt.Println(err)
	}
	return  nil
}