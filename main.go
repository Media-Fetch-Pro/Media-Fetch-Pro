package main

import (
	"github.com/CorrectRoadH/video-tools-for-nas/server"
	_store "github.com/CorrectRoadH/video-tools-for-nas/store"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)
var Store = _store.New()

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/video", server.DownloadVideo)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
