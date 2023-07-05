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
	Id                string   `json:"id"`
	Title             string   `json:"title"`
	Url               string   `json:"url"`
	Status            string   `json:"status"` // unstart, fetching, pending, downloading, finished, failed
	Percent           int      `json:"percent"`
	Size              int      `json:"size"`
	Type              string   `json:"type"`     // video, playlist
	Children          []string `json:"children"` // videos id of playlist
	Author            string   `json:"author"`
	Source            string   `json:"source"`              // bilibili, youtube
	Content           string   `json:"content"`             // the content of the video
	Episode           string   `json:"episode"`             // only of video of playlist
	Parent            string   `json:"parent"`              // the playlist id of the video
	Length            string   `json:"length"`              // the length of the playlist
	StartDownloadTime int64    `json:"start_download_time"` // unix timestamp
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
		go s.Store.SchedulerDownload()

		c.JSON(http.StatusOK, composeResponse("start downloading"))
	})

	g.GET("/videos", func(c *gin.Context) {
		c.JSON(http.StatusOK, composeResponse(s.Store.GetAllVideoInfo()))
	}) // I think may sync all video is not a good idea

	// this is wait to call by python
	g.POST("/update", func(c *gin.Context) {
		var input UpdateVideoStatusInput

		if err := c.ShouldBindJSON(&input); err != nil {
			fmt.Println("update err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.Store.UpdateVideoInfo(types.VideoInfo(input))

		if err != nil {
			c.JSON(http.StatusBadRequest, composeResponse(err.Error()))
			return
		}

		go s.Store.SchedulerDownload()

		c.JSON(http.StatusOK, composeResponse("update video status"))
	})
}
