package store

import (
	"time"

	"github.com/CorrectRoadH/video-tools-for-nas/backend/utils"
)

func (s *Store) DownloadComplete(id string) {
	s.VideosInfo[id].Status = "complete"
	if s.VideosInfo[id].Type == "playlist" {
		for _, child := range s.VideosInfo[id].Children {
			s.VideosInfo[child].Status = "complete"
		}
	}
	// to sleep 2s
	// because the file is not ready to rename
	time.Sleep(2 * time.Second)

	err := utils.RenameVideo(s.VideosInfo[id], s.SystemSettings.StoragePath)
	if err != nil {
		s.VideosInfo[id].Status = "failed"
		if s.VideosInfo[id].Type == "playlist" {
			for _, child := range s.VideosInfo[id].Children {
				s.VideosInfo[child].Status = "failed"
			}
		}

		return
	}
}

func (s *Store) SchedulerDownload() {
	s.schedulerLock.Lock()

	if s.DownloadingVideoNum < s.SystemSettings.MaxDownloadNum {
		for _, value := range s.VideosInfo {

			if value.Status == "complete" {
				continue
			}
			if value.Status == "finished" { // 这里有问题。没有处理好就rename
				if value.Type == "video" {
					s.DownloadComplete(value.Id)
				}

				if value.Type == "playlist" {
					// to check all episodes of playlist
					allFinished := true
					for _, child := range value.Children {
						allFinished = allFinished && (s.VideosInfo[child].Status == "finished")
					}
					if allFinished {
						s.DownloadComplete(value.Id)
					}
				}

				if value.Type == "episode" {
					// to check all episodes of playlist, if all is finished, then set playlist to finished
					allFinished := true
					parentId := value.Parent
					for _, child := range s.VideosInfo[parentId].Children {
						allFinished = allFinished && (s.VideosInfo[child].Status == "finished")
					}
					if allFinished {
						s.VideosInfo[parentId].Status = "finished"
					}
				}
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

			if value.Status == "pending" {
				value.Status = "downloading"
				s.DownloadingVideoNum++
				go utils.DownloadVideo(value, s.SystemSettings.StoragePath)
			}
		}
	}
	s.SaveGlobalVideoInfo()
	s.schedulerLock.Unlock()
}
