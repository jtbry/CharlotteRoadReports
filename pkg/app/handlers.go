package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"actives": s.incidentService.FindActiveIncidents(),
		})
	}
}
