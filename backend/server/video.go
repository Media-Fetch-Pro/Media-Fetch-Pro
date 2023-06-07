package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DownloadVideoInput struct {
	Url     string `json:"url" binding:"required"`
	Storage string `json:"storage" binding:"required"`
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

	c.JSON(http.StatusOK, composeResponse("hello"))
}
