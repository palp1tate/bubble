package routers

import (
	"bubble/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")
	r.GET("/", controllers.IndexHandler)

	v1Group := r.Group("/v1")
	{
		v1Group.POST("/todo", controllers.CreateTodo)

		v1Group.GET("todo/", controllers.GetTodo)

		v1Group.PUT("todo/:id", controllers.UpdateTodo)

		v1Group.DELETE("todo/:id", controllers.DeleteTodo)
	}

	return r
}
