package routers

import (
	"app/app/Controllers/Api/v1"
	"github.com/gin-gonic/gin"
)

func RouterGroupInit(r *gin.Engine)  {
	RouterV1 := r.Group("/v1")
	{
		RouterV1.POST("/uploadLocation",v1.DriverTourTrailController{}.UploadLocation)
		RouterV1.POST("/getLocation",v1.DriverTourTrailController{}.GetLocation)
	}

}
