package server

import "github.com/CorrectRoadH/video-tools-for-nas/store"

var (
	maxDownloadNum = 5
	maxSpeed       = 100
)

func schedulerDownload() {
	currentDownloadNum := 0
	for i := 0; i < store.VideoStatusMap.Len(); i++ {
		if currentDownloadNum >= maxDownloadNum {
			break
		}
		if store.VideoStatusMap[i].Status == "pending" {
			currentDownloadNum++
			go downloadVideo(store.VideoStatusMap[i])
		}
	}
}
