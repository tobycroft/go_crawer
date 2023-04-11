package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz/Input"
)

func IndexController(route *gin.RouterGroup) {

	route.Any("index", index)
	route.Any("login", loginss)
	route.Any("upload", upload)
	route.Any("register")
}

func index(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func loginss(c *gin.Context) {
	password := c.Query("password")
	username := c.Query("username")
	json := map[string]string{}
	json["username"] = username
	json["password"] = password
	gorose.Open()
	c.JSON(0, json)
}

func upload(c *gin.Context) {
	file, ok := Input.Upload(c)
	if !ok {
		return
	}
	fmt.Println(file)
}
