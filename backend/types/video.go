package types

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

type VideoInfo struct {
	Id                string
	Title             string
	Url               string
	Status            string // unstart, fetching, pending, downloading, done, failed
	Percent           int
	Size              int
	Type              string   // video, playlist
	Children          []string // videos id of playlist
	Author            string
	Source            string // bilibili, youtube
	Content           string // the content of the video
	Episode           string // only of video of playlist
	Parent            string // the playlist id of the video
	Length            string // the length of the playlist
	StartDownloadTime int64  // unix timestamp
}

type ConnectionStatus struct {
	BilibiliStatus bool `json:"bilibili_status"`
	YoutubeStatus  bool `json:"youtube_status"`
}
