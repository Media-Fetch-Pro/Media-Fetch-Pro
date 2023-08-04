package server

import "github.com/gin-gonic/gin"

func (s *Server) registerSettingsRoutes(g *gin.RouterGroup) {
	g.GET("/settings", func(c *gin.Context) {
		// get all settings
	})

	g.POST("/settings", func(c *gin.Context) {
		// update settings
	})
}
