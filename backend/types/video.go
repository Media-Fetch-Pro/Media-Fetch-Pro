package types

type VideoInfo struct {
	Id                string `json:"id"`
	Title             string `json:"title"`
	Url               string `json:"url"`
	Status            string `json:"status"` // unstart, fetching, pending, downloading, finished, failed
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
