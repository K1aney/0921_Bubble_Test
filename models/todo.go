package models

import "0921_Bubble/dao"

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo 增删改查
*/

func InitModel() (err error) {
	if err = dao.Db.AutoMigrate(&Todo{}).Error; err != nil {
		return err
	}
	return
}

// CreateTodo 创建todo
func CreateTodo(todo *Todo) (err error) {
	err = dao.Db.Create(&todo).Error
	return err
}

// GetTodoList 查看列表
func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.Db.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodoById(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.Db.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateTodo 更新

func UpdateTodo(todo *Todo) (err error) {
	err = dao.Db.Save(todo).Error
	return
}

// DeleteTodoById 删除

func DeleteTodoById(id string) (err error) {
	err = dao.Db.Where("id=?", id).Delete(&Todo{}).Error
	return
}
