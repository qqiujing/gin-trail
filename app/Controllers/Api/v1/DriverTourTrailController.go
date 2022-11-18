package v1

import (
	"app/app/Controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DriverTourTrailController struct {
	Controllers.BaseController
}

func (con DriverTourTrailController) UploadLocation(c *gin.Context)  {

	data := gin.H{
		"company_id" :   c.Query("company_id"),
		"tour_no"    :   c.Query("tour_no"),
		"driver_id " :   c.Query("driver_id"),
		"lon"        :   c.Query("lon"),
		"lat"        :   c.Query("lat"),
	}

	c.JSON(http.StatusOK, data)

}

func (con DriverTourTrailController)GetLocation(c *gin.Context)  {

}
