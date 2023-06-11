package store

import (
	"fmt"

	"github.com/CorrectRoadH/video-tools-for-nas/utils"
)

// func schedulerDownload() {
// 	currentDownloadNum := 0
// 	// I think may have a additional variable to count the number of downloading videos
// 	for key, value := range store.VideoStatusMap {

// 	}
// }

func (g GlobalVideoStatus) DownloadComplete() {
	g.DownloadingVideoNum--
}

// will be tiger when a video be append to queue and a task is complete
func (g GlobalVideoStatus) SchedulerDownload() {
	if g.DownloadingVideoNum < SystemSettingCache.maxDownloadNum {
		for _, value := range g.VideoStatusMap {
			if value.Status == "pending" {
				value.Status = "fetching"
				g.DownloadingVideoNum++
				// 开始下载
				go utils.DownloadVideo(value, SystemSettingCache.StoragePath)
			} else {
				fmt.Printf(value.Url)
				fmt.Printf("video status is not pending\n")
			}
		}
	} else {
		fmt.Printf("downloading num is max\n")
	}
}
