package controllers

import (
	"encoding/json"
	"fmt"
	"onepiece/database"
	"onepiece/models"
	"onepiece/packages/app"
	"onepiece/packages/e"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Findstockoutlist(c *gin.Context) {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	searchOrdernum := ""
	// searchDetime := ""
	if arg := c.Query("searchOrdernum"); arg != "" {
		searchOrdernum = arg
	}
	// if arg := c.Query("searchDetime"); arg != "" {
	// 	searchDetime = arg
	// }
	testex := map[string]interface{}{
		"page":           page,
		"limit":          limit,
		"searchOrdernum": searchOrdernum,
		// "searchDetime":   searchDetime,
	}
	err, info, total := database.FindOutstock(testex)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total}, "查询成功")
}

func DelNewstockout(c *gin.Context) {
	id := ""
	if arg := c.Query("ordernum"); arg != "" {
		id = arg
	}
	if id == "" {
		app.INFO(c, 30001, "参数错误")
		return
	}
	println(id)
	err := database.DeleteOutstock(id)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "删除成功")
}

func UpdateNewstockout(c *gin.Context) {
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	if m["Ordernum"] == "" {
		app.INFO(c, 30000, "参数非法")
		return
	}

	Ordernum := m["Ordernum"]
	Outpieces := m["Outpieces"]
	Detime := m["Detime"]
	Obton := m["Obton"]
	Ladingnum := m["Ladingnum"]
	// Storagefee := 0.0

	err := database.UpdateOutstock(models.Newstockout{Ordernum: Ordernum, Outpieces: Outpieces, Detime: Detime, Obton: Obton, Ladingnum: Ladingnum})
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "更新成功")
}

func AddNewstockout(c *gin.Context) {
	var stockoutData models.Newstockout
	_ = c.ShouldBindJSON(&stockoutData)

	err := database.AddOutstock(stockoutData)
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	} else {
		app.OK(c, stockoutData, "添加成功")
	}
}

func GetDate2(format string, date1Str string, date2Str string) int {
	// 转化字符串为Time格式
	date1, err := time.ParseInLocation(format, date1Str, time.Local)
	if err != nil {
		return 0
	}
	// 转化字符串为Time格式
	date2, err := time.ParseInLocation(format, date2Str, time.Local)
	if err != nil {
		return 0
	}
	//计算相差天数
	return int(date1.Sub(date2).Hours() / 24)
}

func GetOutstockList(c *gin.Context) {

	var stockinsum int = 0
	var stockoutsum int = 0
	var feesum float64 = 0

	var testsum models.OutstockSearch
	_ = c.ShouldBindQuery(&testsum)
	list, total, _ := database.GetAllstockoutlist(testsum)
	fmt.Println(list)

	var testout models.OutstockSearch
	_ = c.ShouldBindQuery(&testout)
	stockoutData, outsum, _ := database.GetAllstockoutstruct(testout)

	var testin models.InstockSearch
	_ = c.ShouldBindQuery(&testin)
	stockinData, insum, _ := database.GetAllstockinstruct(testin)

	fmt.Println(outsum)
	fmt.Println(insum)
	var i int64 = 0
	var e int64 = 0

	for i < outsum {

		if stockinsum > 0 {
			obton := com.StrTo(stockoutData[i].Obton).MustInt()
			if stockinsum > obton {
				stockinsum = stockinsum - obton
				totaltime := GetDate2("20060102", stockoutData[i].Detime, stockinData[e].Whtime) + 1

				if totaltime > 14 {
					totaltime = totaltime - 14
					var totalcost float64 = float64(totaltime) * float64(obton) * 0.5
					totalstring := totalcost
					stockoutData[i].Storagefee = totalstring
				} else {
					stockoutData[i].Storagefee = 0
				}
				i++

			} else if stockinsum < obton {
				stockoutsum = obton - stockinsum
				totaltime := GetDate2("20060102", stockoutData[i].Detime, stockinData[e].Whtime) + 1
				if totaltime > 14 {
					totaltime = totaltime - 14
					var totalcost float64 = float64(totaltime) * float64(stockinsum) * 0.5
					feesum = feesum + totalcost
				} else {
					feesum = 0
				}
				stockinsum = 0
				e++
			} else {
				stockinsum = 0
				totaltime := GetDate2("20060102", stockoutData[i].Detime, stockinData[e].Whtime) + 1
				if totaltime > 14 {

					totaltime = totaltime - 14
					var totalcost float64 = float64(totaltime) * float64(obton) * 0.5
					totalstring := totalcost
					stockoutData[i].Storagefee = totalstring
				} else {
					stockoutData[i].Storagefee = 0
				}
				i++
				e++
			}
		} else if stockoutsum > 0 {
			ibton := com.StrTo(stockinData[e].Ibton).MustInt()
			if stockoutsum > ibton {

				stockoutsum = stockoutsum - ibton
				totaltime := GetDate2("20060102", stockoutData[i].Detime, stockinData[e].Whtime) + 1
				if totaltime <= 14 {

				} else {
					totaltime = totaltime - 14
					var totalcost float64 = float64(totaltime) * float64(ibton) * 0.5
					feesum = feesum + totalcost
				}
				e++
			} else if stockoutsum < ibton {
				stockinsum = ibton - stockoutsum
				totaltime := GetDate2("20060102", stockoutData[i].Detime, stockinData[e].Whtime) + 1
				if totaltime <= 14 {
					totalstring := feesum
					stockoutData[i].Storagefee = totalstring
					feesum = 0
				} else {
					totaltime = totaltime - 14
					var totalcost float64 = float64(totaltime) * float64(stockoutsum) * 0.5
					totalcost = totalcost + feesum
					totalstring := totalcost
					stockoutData[i].Storagefee = totalstring
					feesum = 0
				}
				stockoutsum = 0
				i++
			} else {
				totaltime := GetDate2("20060102", stockoutData[i].Detime, stockinData[e].Whtime) + 1
				if totaltime <= 14 {
					totalstring := feesum
					stockoutData[i].Storagefee = totalstring
					feesum = 0
				} else {
					totaltime = totaltime - 14
					var totalcost float64 = float64(totaltime) * float64(stockoutsum) * 0.5
					totalcost = totalcost + feesum
					totalstring := totalcost
					stockoutData[i].Storagefee = totalstring
					feesum = 0
				}
				stockoutsum = 0
				i++
				e++
			}
		} else {
			obton := com.StrTo(stockoutData[i].Obton).MustInt()
			ibton := com.StrTo(stockinData[e].Ibton).MustInt()
			if obton > ibton {
				stockoutsum = obton - ibton
				totaltime := GetDate2("20060102", stockoutData[i].Detime, stockinData[e].Whtime) + 1
				if totaltime <= 14 {

				} else {
					totaltime = totaltime - 14
					var totalcost float64 = float64(totaltime) * float64(ibton) * 0.5
					feesum = feesum + totalcost
				}
				e++
			} else if obton < ibton {
				stockinsum = ibton - obton
				totaltime := GetDate2("20060102", stockoutData[i].Detime, stockinData[e].Whtime) + 1
				if totaltime <= 14 {
					stockoutData[i].Storagefee = 0
				} else {
					totaltime = totaltime - 14
					var totalcost float64 = float64(totaltime) * float64(obton) * 0.5
					totalstring := totalcost
					stockoutData[i].Storagefee = totalstring
				}
				i++
			} else {
				totaltime := GetDate2("20060102", stockoutData[i].Detime, stockinData[e].Whtime) + 1
				if totaltime <= 14 {
					stockoutData[i].Storagefee = 0
				} else {
					totaltime = totaltime - 14
					var totalcost float64 = float64(totaltime) * float64(obton) * 0.5
					totalstring := totalcost
					stockoutData[i].Storagefee = totalstring
				}
				i++
				e++
			}
		}
	}
	list = stockoutData
	fmt.Println(stockoutData)
	d := 0
	for d < len(stockoutData) {
		database.Getdatabase().Table("newstockouts").Where("Ordernum = ? And Detime = ?", stockoutData[d].Ordernum, stockoutData[d].Detime).Updates(&stockoutData[d])
		d++
	}
	app.OK(c, map[string]interface{}{"value": list, "total": total}, "获取成功")

}
