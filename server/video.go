package server

import (
	"fmt"
	"net/http"
	"os/exec"

	store "github.com/CorrectRoadH/video-tools-for-nas/store"
	"github.com/CorrectRoadH/video-tools-for-nas/utils"
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

	args := []string{"script/main.py", "--url", input.Url, "--storage", input.Storage, "--type", videoType}
	out, err := exec.Command("python", args...).Output()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, composeResponse(out))
}

// 这个是等 python 来调，来更新状态
func UpdateVideoStatus(c *gin.Context) {

	var input UpdateVideoStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store.AddVideo(store.VideoStatus{
		Id:      input.Url, // 这个id通过url来取个hash值比较好
		Status:  input.Status,
		Percent: input.Percent,
	})
	println(fmt.Print(store.VideoStatusMap))
	c.JSON(http.StatusOK, composeResponse("update video status"))
}

func GetAllVideoStatus(c *gin.Context) {
	c.JSON(http.StatusOK, composeResponse(store.VideoStatusMap))
}
