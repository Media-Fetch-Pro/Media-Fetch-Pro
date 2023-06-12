package server

import (
	"fmt"
	"net/http"

	store "github.com/CorrectRoadH/video-tools-for-nas/backend/store"
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

func DownloadVideo(c *gin.Context) {
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

	// 这里要不要用 goroutine 来做呢?
	store.GlobalVideoStatusMap.AddVideo(types.VideoStatus{
		Id:     utils.GenerateVideoIdFromURL(input.Url),
		Url:    input.Url,
		Status: "pending",
		Type:   videoType,
	})

	store.GlobalVideoStatusMap.SchedulerDownload()
	c.JSON(http.StatusOK, composeResponse("start downloading"))
}

// 这个是等 python 来调，来更新状态
func UpdateVideoStatus(c *gin.Context) {
	var input UpdateVideoStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: I think I will delete some field. because it didn't change when downloading
	store.GlobalVideoStatusMap.UpdateVideoStatus(types.VideoStatus{
		Id:      input.Id, // 这个id通过url来取个hash值比较好
		Title:   input.Title,
		Url:     input.Url,
		Status:  input.Status,
		Percent: input.Percent,
		Size:    input.Size,
	})
	if input.Percent == 100 {
		store.GlobalVideoStatusMap.DownloadComplete()
		fmt.Printf("download complete\n")
		store.GlobalVideoStatusMap.SchedulerDownload()
	}
	c.JSON(http.StatusOK, composeResponse("update video status"))
}

func GetAllVideoStatus(c *gin.Context) {
	c.JSON(http.StatusOK, composeResponse(store.GlobalVideoStatusMap.VideoStatusMap))
}

func GetWebsiteStatus(c *gin.Context) {
	// to check the connection of website: youtube or bilibili. because some people in China can't connect to youtube
	c.JSON(http.StatusOK, composeResponse(""))
}

type ConnectionStatus struct {
	BilibiliStatus bool `json:"bilibili_status"`
	YoutubeStatus  bool `json:"youtube_status"`
}

func GetConnectionStatus(c *gin.Context) {
	// to check the connection of nas
	ConnectionStatus := &ConnectionStatus{
		BilibiliStatus: true,
		YoutubeStatus:  true,
	}
	c.JSON(http.StatusOK, composeResponse(ConnectionStatus))
}
