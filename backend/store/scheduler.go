package store

import (
	"github.com/CorrectRoadH/video-tools-for-nas/backend/utils"
)

// func schedulerDownload() {
// 	currentDownloadNum := 0
// 	// I think may have a additional variable to count the number of downloading videos
// 	for key, value := range store.VideoStatusMap {

// 	}
// }

// func (g GlobalVideoStatus) DownloadComplete() {
// 	g.DownloadingVideoNum--
// }

// // will be tiger when a video be append to queue and a task is complete
// func (g GlobalVideoStatus) SchedulerDownload() {
// 	SaveGlobalVideoStatusMap() // TODO the save timing is not suit
// 	if g.DownloadingVideoNum < SystemSettingCache.maxDownloadNum {
// 		for _, value := range g.VideoStatusMap {
// 			if value.Status == "pending" {
// 				value.Status = "fetching"
// 				g.DownloadingVideoNum++
// 				// 开始下载
// 				go utils.DownloadVideo(value, SystemSettingCache.StoragePath)
// 			} else {
// 				fmt.Printf(value.Url)
// 				fmt.Printf("video status is not pending\n")
// 			}
// 		}
// 	} else {
// 		fmt.Printf("downloading num is max\n")
// 	}
// }

func (s *Store) DownloadComplete(id string) {
	s.VideosInfo[id].Status = "complete"
}

func (s *Store) SchedulerDownload() {
	if s.DownloadingVideoNum < s.SystemSettings.MaxDownloadNum {
		for _, value := range s.VideosInfo {

			// to fetching info
			if value.Status == "unstart" {
				value.Status = "fetching"
				s.DownloadingVideoNum++
				// how to call fetching? sync or async?
				videoInfo, err := utils.FetchingVideoInfo(value)
				if err != nil {
					value.Status = "failed"
				} else {
					value = videoInfo
				}
			}

			// to download video
			// if value.Status == "pending" {
			// 	value.Status = "downloading"
			// 	go utils.DownloadVideo(value, s.SystemSettings.StoragePath)
			// }
		}
	}
}
