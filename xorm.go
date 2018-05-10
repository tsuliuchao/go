package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "pc_dev:123456@tcp(10.112.170.55:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
	}
	//创建表操作
	//	err1 := engine.Sync2(new(User))
	//	fmt.Println(err1)
	//	err2 := engine.Sync2(new(User))
	//	fmt.Println(err2)
	//数据插入操作
	//	user := new(User)
	//	user.Name = "myname"
	//	affected, err := engine.Insert(user)
	//批量插入操作
	users := make([]User, 4)
	users[0].Name = "name1"
	users[1].Name = "name1"
	users[2].Name = "name2"
	users[3].Name = "name3"
	affected, err := engine.Insert(&users)
	fmt.Println(affected, err)
}

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
//https://github.com/go-xorm/xorm
