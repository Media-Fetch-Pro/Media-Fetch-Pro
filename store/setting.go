package store

type SystemSetting struct {
	StoragePath    string
	maxDownloadNum int
	maxSpeed       int
}

var SystemSettingCache = &SystemSetting{
	StoragePath:    "/Users/roadh/Desktop/video-tools-for-nas/video",
	maxDownloadNum: 5,
	maxSpeed:       100,
}
