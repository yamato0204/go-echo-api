package main

import (
	"go-echo-api/controller"
	"go-echo-api/model"
//	"net/http"

	"github.com/labstack/echo/v4"
)


/*
func connect(c echo.Context) error{
	model.Init()
	db, _ := model.DB.DB()
	defer db.Close()
	err := db.Ping()

	if err != nil{
		return c.String(http.StatusInternalServerError, "データベース接続失敗")
	}else{
		return c.String(http.StatusOK, "成功")
	}
}
*/


func main(){
	e := echo.New()

	model.Init()
	db, _ := model.DB.DB()
	defer db.Close()

	
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)


//	e.GET("/", connect)
	e.Logger.Fatal(e.Start(":8080"))
	
}