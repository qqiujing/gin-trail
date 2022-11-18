package Controllers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context)  {
	c.AsciiJSON(200,gin.H{"msg":"hello word"})

}