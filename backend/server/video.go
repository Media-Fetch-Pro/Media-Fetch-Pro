package server

import (
	"encoding/json"
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
	Episode           int      `json:"episode"`             // only of video of playlist
	Parent            string   `json:"parent"`              // the playlist id of the video
	Length            int      `json:"length"`              // the length of the playlist
	StartDownloadTime int64    `json:"start_download_time"` // unix timestamp
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

		// 在单独的goroutine中发送SSE事件
		//go func() {
		//	for {
		//		eventCh <- "data: This is a SSE event\n\n"
		//		time.Sleep(1000)
		//	}
		//}()

		// 循环读取channel中的事件并发送给客户端
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
				// 如果客户端断开连接，则停止发送事件
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

		videoInfo := types.VideoInfo(input)
		err := s.Store.UpdateVideoInfo(videoInfo)
		// TODO move the update to another place.
		// Because it didn't only update place
		s.Store.UpdateChannel <- videoInfo

		if err != nil {
			c.JSON(http.StatusBadRequest, composeResponse(err.Error()))
			return
		}
		go s.Store.SchedulerDownload()

		c.JSON(http.StatusOK, composeResponse("update video status"))
	})
}
