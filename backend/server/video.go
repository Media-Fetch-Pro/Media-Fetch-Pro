package server

import (
	"fmt"
	"net/http"

	"github.com/CorrectRoadH/video-tools-for-nas/backend/types"
	"github.com/CorrectRoadH/video-tools-for-nas/backend/utils"
	"github.com/gin-gonic/gin"
)

type DownloadVideoInput struct {
	Url     string `json:"url" binding:"required"`
	Storage string `json:"storage" binding:"required"`
}

type UpdateVideoStatusInput struct {
	Id      string `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Url     string `json:"url" binding:"required"`
	Status  string `json:"status" binding:"required"`
	Percent int    `json:"percent" binding:"required"`
	Size    int    `json:"size" binding:"required"` // 单位是字节
}

func composeResponse(data any) gin.H {
	return gin.H{"data": data}
}

func (s *Server) registerVideoRoutes(g *gin.RouterGroup) {
	g.GET("/video", func(c *gin.Context) {
		id := c.Param("id")
		videoInfo, err := s.Store.GetVideoInfo(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, composeResponse(err.Error()))
			return
		}
		c.JSON(http.StatusOK, composeResponse(videoInfo))
	})

	g.POST("/video", func(c *gin.Context) {
		var input DownloadVideoInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := utils.HandlerDownloader(input.Url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		videoType := ""
		switch result {
		case "bilibili":
			videoType = "bilibili"
		case "youtube":
			videoType = "youtube"
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "unknown website"})
		}

		s.Store.AddVideoToQueue(types.VideoInfo{
			Id:     utils.GenerateVideoIdFromURL(input.Url),
			Url:    input.Url,
			Status: "unstart",
			Type:   videoType,
		})

		// // 这里要不要用 goroutine 来做呢?
		// store.GlobalVideoStatusMap.AddVideo(types.VideoStatus{
		// 	Id:     utils.GenerateVideoIdFromURL(input.Url),
		// 	Url:    input.Url,
		// 	Status: "pending",
		// 	Type:   videoType,
		// })

		// s.Store.
		// 	store.GlobalVideoStatusMap.SchedulerDownload()
		c.JSON(http.StatusOK, composeResponse("start downloading"))
	})

	g.GET("/videos", func(c *gin.Context) {
		c.JSON(http.StatusOK, composeResponse(s.Store.GetAllVideoInfo()))
	}) // I think may sync all video is not a good idea

	// this is wait to call by python
	g.POST("/update", func(c *gin.Context) {
		var input UpdateVideoStatusInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.Store.UpdateVideoInfo(types.VideoInfo{
			Id:      input.Id,
			Title:   input.Title,
			Url:     input.Url,
			Status:  input.Status,
			Percent: input.Percent,
			Size:    input.Size,
		})

		if err != nil {
			fmt.Println(input)
			c.JSON(http.StatusBadRequest, composeResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, composeResponse("update video status"))
	})
}
