package routers

import (
	"fmt"
	"net/http"
	"onepiece/controllers"
	"onepiece/database"
	"onepiece/models"
	"onepiece/packages/e"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func vesselRouter(r *gin.RouterGroup) {
	r.GET("/vessel", controllers.GetVessel)
	r.DELETE("/vessel", controllers.DelVessel)
	r.PUT("/vessel", controllers.UpdateVessel)
	r.POST("/vessel", controllers.AddVessel)
}

func loginRouter(r *gin.RouterGroup) {
	r.POST("/login", controllers.Loginfunc)
}

func menuRouter(r *gin.RouterGroup) {
	r.GET("/menu", controllers.GetMenu)
}

func logintoRouter(r *gin.RouterGroup) {
	r.GET("/loginto", controllers.GetLoginto)
	r.DELETE("/loginto", controllers.DelLoginto)
	r.PUT("/loginto", controllers.UpdateLoginto)
	r.POST("/loginto", controllers.AddLoginto)
}

func viewRouter(r *gin.RouterGroup) {
	r.GET("/view", GetView)
}

func GetView(c *gin.Context) {

	var testData []models.Newstockout
	err := database.Getdatabase().Table("newstockouts").Select("Ordernum,storsgefee").Find(&testData).Error
	// println(err1)
	fmt.Printf("%+v", testData)
	// err := database.Getdatabase().Table("Test").Find(&testData).Error

	fmt.Printf("testData: %v\n", testData)
	// var d int
	d := len(testData)
	fmt.Printf("%T %v", d, d)
	var y []float64
	fmt.Println()
	var x []string
	for i := 0; i < d; i++ {
		// y[i] = testData[i].Orderfeesum
		y = append(y, testData[i].Storagefee)
		// x[i] = testData[i].Cno
		x = append(x, testData[i].Ordernum)
	}

	if err != nil {
		c.JSON(http.StatusNotFound, models.Result{
			Code:    e.ERROR,
			Message: "查询失败！",
		})
	} else {
		c.JSON(http.StatusOK, models.Result{
			Code:    e.SUCCESS,
			Message: "查询成功！",
			Data: gin.H{
				"x": x,
				"y": y,
			},
		})
	}
}

func view2Router(r *gin.RouterGroup) {
	r.GET("/view2", GetView2)
}

func GetView2(c *gin.Context) {

	var testData []models.Newstockout
	err := database.Getdatabase().Table("newstockouts").Select("Ordernum,obton").Find(&testData).Error
	// println(err1)
	fmt.Printf("%+v", testData)
	// err := database.Getdatabase().Table("Test").Find(&testData).Error

	fmt.Printf("testData: %v\n", testData)
	// var d int
	d := len(testData)
	fmt.Printf("%T %v", d, d)
	var y []int
	fmt.Println()
	var x []string
	for i := 0; i < d; i++ {
		// y[i] = testData[i].Orderfeesum
		y = append(y, com.StrTo(testData[i].Obton).MustInt())
		// com.StrTo(testData[i].Obton).MustInt()
		// x[i] = testData[i].Cno
		x = append(x, testData[i].Ordernum)
	}

	if err != nil {
		c.JSON(http.StatusNotFound, models.Result{
			Code:    e.ERROR,
			Message: "查询失败！",
		})
	} else {
		c.JSON(http.StatusOK, models.Result{
			Code:    e.SUCCESS,
			Message: "查询成功！",
			Data: gin.H{
				"x": x,
				"y": y,
			},
		})
	}
}

func NewstockinRouter(r *gin.RouterGroup) {
	r.GET("/Newstockin", controllers.GetInstockList)
	r.GET("/Newstockins", controllers.FindInstock)
	r.DELETE("/Newstockin", controllers.DeleteInstock)
	r.PUT("/Newstockin", controllers.UpdateInstock)
	r.POST("/Newstockin", controllers.AddInstock)
}

func NewstockoutRouter(r *gin.RouterGroup) {
	r.GET("/Newstockout", controllers.GetOutstockList)
	r.GET("/Newstockouts", controllers.Findstockoutlist)
	r.DELETE("/Newstockout", controllers.DelNewstockout)
	r.PUT("/Newstockout", controllers.UpdateNewstockout)
	r.POST("/Newstockout", controllers.AddNewstockout)
}
