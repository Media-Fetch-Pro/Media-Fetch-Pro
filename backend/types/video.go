package types

type VideoInfo struct {
	// video basic info
	Id      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Url     string `json:"url"`
	Content string `json:"content"` // the content(description) of the video

	// for user
	Status  string `json:"status"` // unstart, fetching, pending, downloading ,finished(download finished), complete, failed
	Reason  string `json:"reason"` // the reason of failed
	Percent int    `json:"percent"`

	Size              int `json:"size"` // the size of the video
	AlreadyDownloaded int `json:"already_downloaded"`

	Type   string `json:"type"`   // video, playlist, episode
	Source string `json:"source"` // bilibili, youtube

	// for playlist
	Parent   string   `json:"parent"`   // the playlist id of the video
	Length   int      `json:"length"`   // the length of the playlist
	Episode  int      `json:"episode"`  // only of video of playlist
	Children []string `json:"children"` // videos id of playlist

	// for download info
	StartDownloadTime int64 `json:"start_download_time"` // unix timestamp
	DownloadSpeed     int64 `json:"download_speed"`      // bytes per second
	EndDownloadTime   int64 `json:"end_download_time"`   // unix timestamp
}

type ConnectionStatus struct {
	BilibiliStatus bool `json:"bilibili_status"`
	YoutubeStatus  bool `json:"youtube_status"`
}
