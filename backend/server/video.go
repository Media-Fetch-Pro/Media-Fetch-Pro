package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/types"
	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/utils"
	"github.com/gin-gonic/gin"
)

type DownloadVideoInput struct {
	Url     string `json:"url" binding:"required"`
	Storage string `json:"storage" binding:"required"`
}

type UpdateVideoStatusInput struct {
	// video basic info
	Id          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Url         string `json:"url"`
	Content     string `json:"content"`      // the content(description) of the video
	PublishTime int64  `json:"publish_time"` // unix timestamp
	Thumbnail   string `json:"thumbnail"`    // the thumbnail of the video
	Tags        string `json:"tags"`         // the tag of the video

	// for user
	Status  string `json:"status"` // unstart, fetching, pending, downloading ,finished(download finished), complete, failed
	Reason  string `json:"reason"` // the reason of failed
	Percent int    `json:"percent"`

	Size              int `json:"size"` // the size of the video
	AlreadyDownloaded int `json:"already_downloaded"`

	Type   string `json:"type"`   // video, playlist, episode
	Source string `json:"source"` // bilibili, youtube

	// for playlist
	Parent   string   `json:"parent"`   // the playlist id of the video
	Length   int      `json:"length"`   // the length of the playlist
	Episode  int      `json:"episode"`  // only of video of playlist
	Children []string `json:"children"` // videos id of playlist

	// for download info
	StartDownloadTime int64 `json:"start_download_time"` // unix timestamp
	DownloadSpeed     int64 `json:"download_speed"`      // bytes per second
	EndDownloadTime   int64 `json:"end_download_time"`   // unix timestamp
}

func composeResponse(data any) gin.H {
	return gin.H{"data": data}
}

func (s *Server) registerVideoRoutes(g *gin.RouterGroup) {
	g.GET("/history", func(c *gin.Context) {
		// the API is a server sent event API
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		for {
			select {
			case event := <-s.Store.UpdateChannel:
				json, err := json.Marshal(event)
				if err != nil {
					fmt.Println("json err: ", err)
				}
				fmt.Fprintf(c.Writer, "data: "+string(json)+"\n\n")
				c.Writer.Flush()
			case <-c.Writer.CloseNotify():
				return
			}
		}
	})

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

		// to process storage
		s.Store.SystemSettings.StoragePath = input.Storage

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

		videoInfo := types.VideoInfo(input)
		err := s.Store.UpdateVideoInfoPartition(videoInfo)

		if err != nil {
			c.JSON(http.StatusBadRequest, composeResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, composeResponse("update video status"))
	})
}
