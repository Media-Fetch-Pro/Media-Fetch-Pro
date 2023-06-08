package main

import (
	"github.com/CorrectRoadH/video-tools-for-nas/server"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/video", server.DownloadVideo)
	r.POST("/update", server.UpdateVideoStatus)
	r.GET("/video", server.GetAllVideoStatus)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
