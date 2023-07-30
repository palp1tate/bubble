package models

import (
	"bubble/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//Todo 增删改查

// CTodo 创建todo
func CTodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

// GTodo 查询todo
func GTodo() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	return
}

func UTodo(id string) (todo *Todo, err error) {
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func STodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
