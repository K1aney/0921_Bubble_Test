package main

import (
	"0921_Bubble/dao"
	"0921_Bubble/models"
	"0921_Bubble/routers"
)

func main() {
	// 创建数据库
	// CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)

	}
	// 程序关闭退出数据库连接
	defer dao.DbClose()
	// 模型绑定
	err1 := models.InitModel()
	if err1 != nil {
		panic(err1)
	}

	r := routers.SetupRouter()

	r.Run(":9090")
}
