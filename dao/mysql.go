package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 初始化全局变量
var (
	Db *gorm.DB
)

// go语言通过首字母是否大小写，来表示是否对外可建
func InitMySql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = Db.DB().Ping()
	return
}

func DbClose() {
	// 程序退出，退出数据库
	Db.Close()
}
