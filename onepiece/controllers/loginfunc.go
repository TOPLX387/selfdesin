package controllers

import (
	"encoding/json"
	"onepiece/function"
	"onepiece/packages/app"
	"onepiece/packages/e"

	"github.com/gin-gonic/gin"
)

func Loginfunc(c *gin.Context) {
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	username := m["username"]
	password := m["password"]
	err, logintoData, _ := function.GetLogintoByName(username)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	if len(logintoData) < 1 {
		app.INFO(c, 30000, "没有此用户")
		return
	}
	if logintoData[0].LogintoPassword != password {
		app.INFO(c, 30001, "密码错误")
		return
	}
	app.OK(c, map[string]interface{}{"user": logintoData[0]}, "登录成功")
}
