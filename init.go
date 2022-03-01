package ginXm

import (
	"ginXm/router"
	"github.com/gin-gonic/gin"
)

func SystemInit() {
	r := gin.Default()
	router.SystemRouter(r)
	r.Run(":8088")
}
