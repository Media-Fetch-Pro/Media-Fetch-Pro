package store

type SystemSetting struct {
	StoragePath    string
	maxDownloadNum int
	maxSpeed       int
}

var SystemSettingCache = &SystemSetting{
	StoragePath:    "/home/ctrdh/video",
	maxDownloadNum: 5,
	maxSpeed:       100,
}
