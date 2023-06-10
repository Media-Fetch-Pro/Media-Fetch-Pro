package server

import (
	"github.com/CorrectRoadH/video-tools-for-nas/store"
)

var (
	maxDownloadNum = 5
	maxSpeed       = 100
)

func schedulerDownload() {
	currentDownloadNum := 0
	// I think may have a additional variable to count the number of downloading videos
	for key, value := range store.VideoStatusMap {

	}
}
