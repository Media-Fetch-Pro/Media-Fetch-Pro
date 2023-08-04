package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerSettingsRoutes(g *gin.RouterGroup) {
	g.GET("/settings", func(c *gin.Context) {
		// get all settings
	})

	g.POST("/settings", func(c *gin.Context) {
		// update settings
	})

	g.GET("/systemStatus", func(c *gin.Context) {
		// get status
		c.JSON(http.StatusOK, composeResponse(s.Profile))
	})
}
