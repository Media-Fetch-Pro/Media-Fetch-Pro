package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CorrectRoadH/video-tools-for-nas/backend/types"
	"github.com/gin-gonic/gin"
)

func checkWebsiteConnection(url string) bool {
	// to request address param
	// if response code is 200, then return true
	// else return false
	// timeout is 5s
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	_, err := client.Get(url)
	if err == nil {
		return true
	} else {
		fmt.Println(err)
		return false
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
