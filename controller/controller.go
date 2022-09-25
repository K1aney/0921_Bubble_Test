package controller

import (
	"0921_Bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	只调用，不去实现
	url  	=> controller 	=> logic 		=> model
	请求	=>	控制器		=>	业务逻辑		=>模型层的增删改查
*/

// 首页
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 创建
func CreateTodo(c *gin.Context) {
	//	前端页面填写待办事项 => 会发请求到此处
	//	1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	//	2.存入数据库
	err := models.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 查看
func GetTodoList(c *gin.Context) {

	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// 更新
func UpdataTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效id",
		})
		return
	}
	todo, err := models.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
	err = models.UpdateTodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 删除
func DeleteTodo(c *gin.Context) {
	// 获取id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效id",
		})
	}
	// 根据id删数据
	if err := models.DeleteTodoById(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": "删除成功",
		})
	}
}
