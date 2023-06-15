package main

import (
	"net/http"

	"github.com/CorrectRoadH/video-tools-for-nas/backend/server"
	store "github.com/CorrectRoadH/video-tools-for-nas/backend/store"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.NoRoute(gin.WrapH(http.FileServer(server.GetFileSystem("dist"))))

	apiv1 := r.Group("api")
	{
		apiv1.POST("/video", server.DownloadVideo)
		apiv1.POST("/update", server.UpdateVideoStatus)
		apiv1.GET("/video", server.GetAllVideoStatus) // I think may sync all video is not a good idea
		// only update downloading video status
		// apiv1.GET("/video/:id", server.GetVideoStatus)
		apiv1.GET("/status", server.GetConnectionStatus)
		// TODO: delete video
		//
	}
	return r
}

func main() {
	store.LoadGlobalVideoStatusMap()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:7789
	r.Run(":7789")
}
