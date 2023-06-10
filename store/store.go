package store

type Store struct {
	videoStatus string
}

type VideoStatus struct {
	Id                    string
	Title                 string
	Url                   string
	Status                string // pending, downloading, done, failed
	Percent               int
	Size                  int
	Type                  string // bilibili, youtube
	AlreadyDownloadedSize int
	isCollection          bool
	CollectionId          string // 用来标记是哪个 collection 的。比如一个视频是属于另一个视频的集合的，那么这个就是那个集合的 id。
}

type GlobalVideoStatus struct {
	VideoStatusMap      map[string]*VideoStatus
	DownloadingVideoNum int
}

// 这个怎么做持久化比如好呢? 然后要不要不要用全局变量会不会比较好?
var VideoStatusMap map[string]*VideoStatus = make(map[string]*VideoStatus)

func AddVideo(videoStatus VideoStatus) {
	VideoStatusMap[videoStatus.Id] = &videoStatus
}
