package store

type Store struct {
	videoStatus string
}

type VideoStatus struct {
	VideoId string
	Status  string
	Percent int
}

// 这个怎么做持久化比如好呢? 然后要不要不要用全局变量会不会比较好?
var VideoStatusMap map[string]*VideoStatus = make(map[string]*VideoStatus)

func AddVideo(videoStatus VideoStatus) {
	VideoStatusMap[videoStatus.VideoId] = &videoStatus
}
