package main

import (
	"app/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.RouterGroupInit(r)
	r.Run(":8080")
}

/*func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sddd",
		})
	})
	r.POST("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pppp",
		})
	})
	r.PUT("ping", func(ctx *gin.Context) {

	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}*/
