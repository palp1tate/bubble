package controllers

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//url -->controller -->logic -->model
//请求来了 -->控制器 -->业务逻辑 -->模型层的增删改查

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)

}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		panic(err)
	}
	if err := models.CTodo(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todo,
		})
	}
}

func GetTodo(c *gin.Context) {

	if todoList, err := models.GTodo(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.UTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.BindJSON(&todo)
	if err != nil {
		panic(err)
	}
	if err = models.STodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
