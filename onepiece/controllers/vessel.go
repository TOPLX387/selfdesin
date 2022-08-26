package controllers

import (
	"encoding/json"
	"onepiece/database"
	"onepiece/models"
	"onepiece/packages/app"
	"onepiece/packages/e"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetVessel(c *gin.Context) {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	searchLadingnum := ""
	if arg := c.Query("searchLadingnum"); arg != "" {
		searchLadingnum = arg
	}
	testex := map[string]interface{}{
		"page":            page,
		"limit":           limit,
		"searchLadingnum": searchLadingnum,
	}
	err, info, total := database.GetAllVessel(testex)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total}, "查询成功")
}

func DelVessel(c *gin.Context) {
	id := ""
	// fmt.Println("test呵呵哈哈哈或或")

	if arg := c.Query("cname"); arg != "" {
		id = arg

	}
	if id == "" {
		app.INFO(c, 30001, "参数错误")
		return
	}
	println(id)
	err := database.DelVessel(id)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "删除成功")
}

func UpdateVessel(c *gin.Context) {
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	if m["Ladingnum"] == "" {
		app.INFO(c, 30000, "参数非法")
		return
	}
	ladingnum := m["Ladingnum"]
	cname := m["Cname"]
	destination := m["Destination"]
	err := database.UpdateVessel(models.Vessel{Ladingnum: ladingnum, Cname: cname, Destination: destination})
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "更新成功")
}

func AddVessel(c *gin.Context) {
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	ladingnum := m["Ladingnum"]
	cname := m["Cname"]
	destination := m["Destination"]
	err := database.AddVessel(models.Vessel{Ladingnum: ladingnum, Cname: cname, Destination: destination})
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "添加成功")
}
