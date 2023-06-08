package store

type Store struct {
	videoStatus string
}

type VideoStatus struct {
	VideoId string
	Status  string
}

var VideoStatusMap map[string]*VideoStatus = make(map[string]*VideoStatus)

func New() Store {
	return Store{
		videoStatus: "213",
	}
}

func AddVideo(videoStatus VideoStatus) {
	VideoStatusMap[videoStatus.VideoId] = &videoStatus
}

var StoreCache *Store = &Store{
	videoStatus: "123",
}
