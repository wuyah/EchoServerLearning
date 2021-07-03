package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main(){
	e := echo.New()
	//设置访问静态资源
	e.File("/user","resource/abi.png")

	//获取GET信息
	e.GET("/users", getUsers)

	//获取POST信息
	e.POST("/users", getPost)
	e.Start(":8088")
}

func getPost(c echo.Context) error{
	//获取url上的path参数， url模式里面定义了参数：id
	name := c.FormValue("name")

	fmt.Printf("%s", name)

	return c.String(http.StatusOK, name)
}

func getUsers(c echo.Context) error{
	username := c.QueryParam("username")
	return c.String(http.StatusOK, username)

}