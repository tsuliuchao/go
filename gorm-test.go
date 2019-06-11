package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type tbl_class struct {
	id         uint   `gorm:"primary_key"`
	uid        int    // 外键 (属于), tag `index`是为该列创建索引
	class_name string // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	score      float32
}
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	CreatedBy uint
}
type CSDN struct {
	gorm.Model
	AId     string `gorm:default:'fuck'`
	Content string
}

func main() {
	//mysql -upc_dev -p -h10.112.170.55 -upc_dev -p123456
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "pc_dev:123456@tcp(10.112.170.55:3306)/test?charset=utf8&parseTime=true")
	defer db.Close()
	var TBL = tbl_class{}
	aa := db.First(&TBL)
	fmt.Println(aa, err)
	db.HasTable("CSDN")
	db.HasTable(&CSDN{})
	db.SingularTable(true) //全局设置表名不可以为复数形式。
	db.CreateTable(&CSDN{})
	//db.DropTable(&CSDN{})
	c := CSDN{AId: "100"}
	db.Create(&c)
	m := db.NewRecord(c)
	fmt.Println(m)
	//	db.Where("id=?", 1).Delete(&CSDN{})
	var csd CSDN
	db.First(&csd)
	fmt.Println(csd)

}
