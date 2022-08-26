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

func AddInstock(c *gin.Context) {
	var stockinData models.Newstockin
	_ = c.ShouldBindJSON(&stockinData)

	err := database.AddInstock(stockinData)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	} else {
		app.OK(c, stockinData, "添加成功")
	}
}

// func DeleteInstock(c *gin.Context) {
// 	var stockinData models.Newstockin
// 	_ = c.ShouldBindJSON(&stockinData)

// 	err := database.DeleteInstock(stockinData)
// 	if err != nil {
// 		app.Error(c, e.ERROR, err, err.Error())
// 		return
// 	} else {
// 		app.OK(c, stockinData, "删除成功")
// 	}
// }

func DeleteInstock(c *gin.Context) {
	id := ""
	if arg := c.Query("lpnum"); arg != "" {
		id = arg
	}
	if id == "" {
		app.INFO(c, 30001, "参数错误")
		return
	}
	err := database.DeleteInstock(id)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "删除成功")
}

func UpdateInstock(c *gin.Context) {
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	if m["Lpnum"] == "" {
		app.INFO(c, 30000, "参数非法")
		return
	}

	Lpnum := m["Lpnum"]
	Ibpieces := m["Ibpieces"]
	Whtime := m["Whtime"]
	Ibton := m["Ibton"]

	err := database.UpdateInstock(models.Newstockin{Lpnum: Lpnum, Ibpieces: Ibpieces, Whtime: Whtime, Ibton: Ibton})
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "更新成功")
}

func FindInstock(c *gin.Context) {
	// var stockinData models.Newstockin
	// _ = c.ShouldBindQuery(&stockinData)

	// reinstock, err := database.FindInstock(stockinData.Lpnum)
	// if err != nil {
	// 	app.Error(c, e.ERROR, err, err.Error())
	// 	return
	// } else {
	// 	// app.OK(c, gin.H{"reinstock": reinstock}, "查询成功")
	// 	app.OK(c, map[string]interface{}{"value": reinstock}, "查询成功")
	// }
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	searchLpnum := ""
	if arg := c.Query("searchLpnum"); arg != "" {
		searchLpnum = arg
	}
	testex := map[string]interface{}{
		"page":        page,
		"limit":       limit,
		"searchLpnum": searchLpnum,
	}
	err, info, total := database.FindInstock(testex)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total}, "查询成功")
}

func GetInstockList(c *gin.Context) {
	var pageInfo models.InstockSearch
	_ = c.ShouldBindQuery(&pageInfo)
	list, total, _ := database.GetAllstockinstruct(pageInfo)

	app.OK(c, map[string]interface{}{"value": list, "total": total}, "获取成功")
	// c.JSON(http.StatusOK, models.Result{
	// 	Code:    e.SUCCESS,
	// 	Message: "获取成功！",
	// 	Data: models.PageResult{
	// 		List:     list,
	// 		Total:    total,
	// 		Page:     pageInfo.Page,
	// 		PageSize: pageInfo.PageSize,
	// 	},
	// })
}
