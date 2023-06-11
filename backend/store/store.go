package store

import "github.com/CorrectRoadH/video-tools-for-nas/backend/types"

type GlobalVideoStatus struct {
	VideoStatusMap      map[string]*types.VideoStatus
	DownloadingVideoNum int
}

// 这个怎么做持久化比如好呢? 然后要不要不要用全局变量会不会比较好?
var GlobalVideoStatusMap *GlobalVideoStatus = &GlobalVideoStatus{
	VideoStatusMap:      make(map[string]*types.VideoStatus),
	DownloadingVideoNum: 0,
}

// func AddVideo(videoStatus types.VideoStatus) {
// 	VideoStatusMap[videoStatus.Id] = &videoStatus
// }

func (g *GlobalVideoStatus) AddVideo(videoStatus types.VideoStatus) {
	g.VideoStatusMap[videoStatus.Id] = &videoStatus
}

func (g *GlobalVideoStatus) UpdateVideoStatus(videoStatus types.VideoStatus) {
	g.VideoStatusMap[videoStatus.Id] = &videoStatus
}
