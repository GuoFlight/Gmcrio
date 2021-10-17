package v1

import (
	"Gmicro/http/httpCommon"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context){
	httpCommon.HttpLogger.Error("health")
	c.JSON(200, "alive")
}
