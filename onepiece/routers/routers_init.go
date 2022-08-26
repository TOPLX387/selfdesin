package routers

import (
	"onepiece/middleWare"

	"github.com/gin-gonic/gin"
)

func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/route")
	menuRouter(r)
	logintoRouter(r)
	loginRouter(r)
	vesselRouter(r)
	viewRouter(r)
	view2Router(r)
	NewstockinRouter(r)
	NewstockoutRouter(r)

}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleWare.Cors())
	g := r.Group("/guobaichuan")
	sysNoCheckRoleRouter(g)

	return r
}
