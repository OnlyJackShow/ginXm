package ginXm

import (
	"github.com/OnlyJackShow/ginXm/router"
	"github.com/gin-gonic/gin"
)

func SystemInit() *gin.Engine {
	r := gin.Default()
	router.SystemRouter(r)
	//r.Run(":8088")
	return r
}
