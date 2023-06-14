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

	// r.NoRoute("/", server.GetFileSystem("dist"))
	// r.NoRoute(gin.WrapH(http.FileServer(gin.Dir("static", false))))
	r.StaticFS("/ui", server.GetFileSystem("dist"))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ui")
		c.Abort()
	})

	apiv1 := r.Group("api")
	{
		apiv1.POST("/video", server.DownloadVideo)
		apiv1.POST("/update", server.UpdateVideoStatus)
		apiv1.GET("/video", server.GetAllVideoStatus)
		apiv1.GET("/status", server.GetConnectionStatus)
	}
	return r
}

func main() {
	store.LoadGlobalVideoStatusMap()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
