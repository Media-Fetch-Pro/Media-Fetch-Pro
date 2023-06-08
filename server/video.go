package server

import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"

	store "github.com/CorrectRoadH/video-tools-for-nas/store"
	"github.com/gin-gonic/gin"
)

type DownloadVideoInput struct {
	Url     string `json:"url" binding:"required"`
	Storage string `json:"storage" binding:"required"`
}

type UpdateVideoStatusInput struct {
	Url    string `json:"url" binding:"required"`
	Status string `json:"status" binding:"required"`
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

	result, err := HandlerDownloader(input.Url)
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

func UpdateVideoStatus(c *gin.Context) {
	println("update video status")
	println(fmt.Print(store.VideoStatusMap))
	store.AddVideo(store.VideoStatus{
		VideoId: "123",
		Status:  "123",
	})
	println(fmt.Print(store.VideoStatusMap))
	c.JSON(http.StatusOK, composeResponse("update video status"))
}

func GetAllVideoStatus(c *gin.Context) {
	c.JSON(http.StatusOK, composeResponse(store.VideoStatusMap))
}

func HandlerDownloader(url string) (string, error) {
	// 判断url是bilibili还是youtube
	// 如果是bilibili，则返回bilibili
	// 如果是youtube，则返回youtube
	// 如果是其他，返回错误
	// 这里用正则做判断，如 https://www.youtube.com/watch?v=KO-yFbjtXGg
	// https://www.bilibili.com/video/BV1Y7411t7zZ
	youtubeRegex := regexp.MustCompile(`https?://www.youtube.com/watch\?v=.*`)
	bilibiliRegex := regexp.MustCompile(`https?://www.bilibili.com/video/.*`)

	if youtubeRegex.MatchString(url) {
		return "youtube", nil
	} else if bilibiliRegex.MatchString(url) {
		return "bilibili", nil
	} else {
		return "unknown", fmt.Errorf("unknown website")
	}

}
