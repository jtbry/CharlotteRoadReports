package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jtbry/CharlotteRoadReports/models"
)

func activeIncidents(ctx *gin.Context) {
	ctx.JSON(200, models.FindActiveIncidents())
}

// Register endpoints for the /api/incidents route
func RegisterIncidentsApi(router *gin.RouterGroup) {
	router.GET("/incidents/active", activeIncidents)
}
