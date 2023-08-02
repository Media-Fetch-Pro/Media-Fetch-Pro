package store

import (
	"fmt"
	"time"

	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/types"
	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/utils"
)

func (s *Store) DownloadComplete(id string) {
	s.UpdateVideoInfoPartition(types.VideoInfo{
		Id:      id,
		Status:  "complete",
		Percent: 100,
	})

	if s.VideosInfo[id].Type == "playlist" {
		for _, childId := range s.VideosInfo[id].Children {
			s.UpdateVideoInfoPartition(types.VideoInfo{
				Id:     childId,
				Status: "complete",
			})
		}
	}

	// to sleep 2s
	// because the file is not ready to rename
	time.Sleep(2 * time.Second)

	err := utils.RenameVideo(s.VideosInfo[id], s.SystemSettings.StoragePath)
	if err != nil {
		s.UpdateVideoInfoPartition(types.VideoInfo{
			Id:     id,
			Status: "failed",
		})

		if s.VideosInfo[id].Type == "playlist" {
			for _, childId := range s.VideosInfo[id].Children {
				s.UpdateVideoInfoPartition(types.VideoInfo{
					Id:     childId,
					Status: "failed",
				})
			}
		}
	}
}

func (s *Store) SchedulerDownload() {
	// I think this is not a good way to do this🤔
	s.SchedulerLock.Lock()
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
					downloadFinishNum := 0
					for _, child := range s.VideosInfo[parentId].Children {
						allFinished = allFinished && (s.VideosInfo[child].Status == "finished")
						if s.VideosInfo[child].Status == "finished" {
							downloadFinishNum++
						}
					}

					if allFinished {
						s.VideosInfo[parentId].Status = "finished"
					}

					// update playlist percent
					s.UpdateVideoInfoPartition(types.VideoInfo{
						Id:      parentId,
						Percent: int(float64(downloadFinishNum) / float64(len(s.VideosInfo[parentId].Children)) * 100),
					})
				}
			}

			// to fetching info
			if value.Status == "unstart" {
				s.DownloadingVideoNum++
				// how to call fetching? sync or async?
				fmt.Println("fetch start")
				err := utils.FetchingVideoInfo(value)
				if err != nil {
					s.UpdateVideoInfoPartition(types.VideoInfo{
						Id:     value.Id,
						Status: "failed",
					})
				}
			}

			if value.Status == "pending" {
				value.Status = "downloading"
				s.DownloadingVideoNum++
				go func(value *types.VideoInfo) {
					err := utils.DownloadVideo(value, s.SystemSettings.StoragePath)
					if err != nil {
						s.UpdateVideoInfoPartition(types.VideoInfo{
							Id:     value.Id,
							Status: "failed",
						})
					}
				}(value)
			}
		}
	}

	s.SaveGlobalVideoInfo()
	s.SchedulerLock.Unlock()
}
