package controllers

import (
	"encoding/json"
	"onepiece/database"
	"onepiece/models"
	"onepiece/packages/app"
	"onepiece/packages/e"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetLoginto(c *gin.Context) {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	searchName := ""
	if arg := c.Query("searchName"); arg != "" {
		searchName = arg
	}
	logintoParam := map[string]interface{}{
		"page":       page,
		"limit":      limit,
		"searchName": searchName,
	}
	err, info, total := database.GetAllLoginto(logintoParam)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total}, "查询成功")
}

func DelLoginto(c *gin.Context) {
	id := -1
	if arg := c.Query("id"); arg != "" {
		id = com.StrTo(arg).MustInt()
	}
	if id == -1 {
		app.INFO(c, 30001, "参数错误")
		return
	}
	err2, loginto := database.GetLogintoById(id)
	if err2 != nil {
		app.Error(c, e.ERROR, err2, err2.Error())
		return
	}
	if loginto.LogintoLevel == 2 {
		app.INFO(c, 30000, "管理员用户不能删除")
		return
	}
	err := database.DelLoginto(id)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "删除成功")
}

func UpdateLoginto(c *gin.Context) {
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	if m["id"] == "" {
		app.INFO(c, 30000, "参数非法")
		return
	}
	id := -1
	id = com.StrTo(m["id"]).MustInt()
	logintoName := m["logintoName"]
	logintoPassword := m["logintoPassword"]
	logintoLevel := 0
	logintoLevel = com.StrTo(m["logintoLevel"]).MustInt()
	logintoRemarks := m["logintoRemarks"]
	err := database.UpdateLoginto(models.Loginto{LogintoId: id, LogintoName: logintoName, LogintoPassword: logintoPassword, LogintoLevel: logintoLevel, LogintoRemarks: logintoRemarks, UpdateTime: time.Now()})
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "更新成功")
}

func AddLoginto(c *gin.Context) {
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	logintoName := m["logintoName"]
	logintoPassword := m["logintoPassword"]
	logintoLevel := com.StrTo(m["logintoLevel"]).MustInt()
	logintoRemarks := m["logintoRemarks"]
	err := database.AddLoginto(models.Loginto{LogintoId: -1, LogintoName: logintoName, LogintoPassword: logintoPassword, LogintoLevel: logintoLevel, LogintoRemarks: logintoRemarks, UpdateTime: time.Now()})
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "添加成功")
}
