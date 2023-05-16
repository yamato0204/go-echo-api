package controller

import (
	"go-echo-api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)


func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil{
		return err
	}
	model.DB.Create(&user)

	return c.JSON(http.StatusCreated, user)
}


func GetUsers(c echo.Context)error{
	users := []model.User{}
	
	model.DB.Find(&users)
	return c.JSON(http.StatusOK, users)

}



func GetUser(c echo.Context) error{
	user := model.User{}

	if err := c.Bind(&user); err != nil{
		return err 
	}
	model.DB.Take(&user)
	return c.JSON(http.StatusOK, user)

}


func UpdateUser(c echo.Context) error{
	user := model.User{}
	if err := c.Bind(&user); err != nil{
		return err
	}
	model.DB.Updates(&user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return nil
	}

	model.DB.Delete(&user)
	return c.JSON(http.StatusOK, user)
}



