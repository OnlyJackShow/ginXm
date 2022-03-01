package router

import (
	"github.com/gin-gonic/gin"
)

func SystemRouter(r *gin.Engine) {

	r.NoMethod(func(c *gin.Context) {
		c.JSON(404, "未找到该方法")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, "未找到该路由")
	})
	g := r.Group("/v1/service/")
	{
		g.GET("hello", func(c *gin.Context) {
			c.JSON(200, "hello")
		})
	}

}
