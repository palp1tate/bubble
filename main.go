package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {

	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}

	err = dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}
	r := routers.SetupRouter()
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
