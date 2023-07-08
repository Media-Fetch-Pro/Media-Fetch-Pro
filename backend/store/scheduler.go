package store

import (
	"fmt"

	"github.com/CorrectRoadH/video-tools-for-nas/backend/utils"
)

func (s *Store) DownloadComplete(id string) {
	s.VideosInfo[id].Status = "complete"
}

func (s *Store) SchedulerDownload() {
	fmt.Println(s.VideosInfo)

	if s.DownloadingVideoNum < s.SystemSettings.MaxDownloadNum {
		for _, value := range s.VideosInfo {

			if value.Status == "finished" {
				continue
			}

			// to fetching info
			if value.Status == "unstart" {
				value.Status = "fetching"
				s.DownloadingVideoNum++
				// how to call fetching? sync or async?
				err := utils.FetchingVideoInfo(value)
				if err != nil {
					value.Status = "failed"
				}
			}

			if value.Status == "pending" && value.Type != "playlist" {
				value.Status = "downloading"
				s.DownloadingVideoNum++
				go utils.DownloadVideo(value, s.SystemSettings.StoragePath)
			}
		}
	}
}
