package routers

import (
	"0921_Bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")

	// 引入前端页面，告诉gin框架找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	// api(v1)
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看
		v1Group.GET("/todo", controller.GetTodoList)
		// 更新
		v1Group.PUT("/todo/:id", controller.UpdataTodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
