package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (con BaseController) Success(c *gin.Context) {
	c.String(http.StatusOK, "successssssss")
	
}
