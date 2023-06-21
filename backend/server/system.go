package server

import (
	"net/http"

	"github.com/CorrectRoadH/video-tools-for-nas/backend/types"
	"github.com/gin-gonic/gin"
)

func checkWebsiteConnection(address string) bool {
	// to request address param
	// if response code is 200, then return true
	// else return false
	_, err := http.Get(address)
	if err == nil {
		return false
	} else {
		return true
	}
}

func (s *Server) registerSystemRoutes(g *gin.RouterGroup) {
	g.GET("/status", func(c *gin.Context) {
		// to check the connection of nas
		ConnectionStatus := &types.ConnectionStatus{
			BilibiliStatus: checkWebsiteConnection("https://www.bilibili.com/"),
			YoutubeStatus:  checkWebsiteConnection("https://www.youtube.com/"),
		}
		// to check the connection of website: youtube or bilibili. because some people in China can't connect to youtube
		c.JSON(http.StatusOK, composeResponse(ConnectionStatus))
	})
}
