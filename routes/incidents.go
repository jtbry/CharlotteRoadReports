package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jtbry/CharlotteRoadReports/models"
)

func activeIncidents(ctx *gin.Context) {
	actives := make([]models.Incident, 0)
	models.DB.Where("is_active = ?", 1).Find(&actives)
	ctx.JSON(200, actives)
}

func RegisterIncidentsApi(router *gin.RouterGroup) {
	router.GET("/incidents/active", activeIncidents)
}
