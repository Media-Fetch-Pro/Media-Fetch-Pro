package types

type VideoInfo struct {
	Id                string   `json:"id"`
	Title             string   `json:"title"`
	Url               string   `json:"url"`
	Status            string   `json:"status"` // unstart, fetching, pending, downloading, finished, failed
	Percent           int      `json:"percent"`
	Size              int      `json:"size"`
	Type              string   `json:"type"`     // video, playlist
	Children          []string `json:"children"` // videos id of playlist
	Author            string   `json:"author"`
	Source            string   `json:"source"`              // bilibili, youtube
	Content           string   `json:"content"`             // the content of the video
	Episode           string   `json:"episode"`             // only of video of playlist
	Parent            string   `json:"parent"`              // the playlist id of the video
	Length            string   `json:"length"`              // the length of the playlist
	StartDownloadTime int64    `json:"start_download_time"` // unix timestamp
}

type ConnectionStatus struct {
	BilibiliStatus bool `json:"bilibili_status"`
	YoutubeStatus  bool `json:"youtube_status"`
}
