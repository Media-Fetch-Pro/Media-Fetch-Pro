package store

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/types"
	"gopkg.in/yaml.v3"
)

type SystemSettings struct {
	StoragePath    string
	MaxDownloadNum int
	MaxSpeed       int
}

type Store struct {
	db *sql.DB

	VideosInfo          map[string]*types.VideoInfo
	DownloadingVideoNum int // the video num that is downloading
	SystemSettings      SystemSettings
	SchedulerLock       sync.Mutex
	UpdateChannel       chan types.VideoInfo
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:             db,
		VideosInfo:     make(map[string]*types.VideoInfo),
		UpdateChannel:  make(chan types.VideoInfo, 1000),
		SystemSettings: LoadSystemSetting(),
	}
}

func (s *Store) GetDB() *sql.DB {
	return s.db
}

func (s *Store) SaveGlobalVideoInfo() {
	// handle err
	cwd, err := os.Getwd()
	// handle err
	if err != nil {
		fmt.Println(err)
	}

	raw, err := yaml.Marshal(s.VideosInfo)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(
		path.Join(cwd, "video.yaml"),
		raw,
		0644,
	)
	if err != nil {
		fmt.Println(err)
	}

}

func (s *Store) LoadGlobalVideoInfo() {
	cwd, err := os.Getwd()
	// handle err
	if err != nil {
		fmt.Println(err)
	}

	raw, err := os.ReadFile(
		path.Join(cwd, "video.yaml"),
	)
	var status map[string]*types.VideoInfo

	if err != nil {
		if os.IsNotExist(err) {
			status = make(map[string]*types.VideoInfo)
		}
	} else {
		yaml.Unmarshal(raw, &status)
	}

	s.VideosInfo = status
}

// type GlobalVideoStatus struct {
// 	VideoStatusMap      map[string]*types.VideoStatus
// 	DownloadingVideoNum int
// }

// // 这个怎么做持久化比如好呢? 然后要不要不要用全局变量会不会比较好?
// var GlobalVideoStatusMap *GlobalVideoStatus = &GlobalVideoStatus{
// 	VideoStatusMap:      make(map[string]*types.VideoStatus),
// 	DownloadingVideoNum: 0,
// }

// func (g *GlobalVideoStatus) AddVideo(videoStatus types.VideoStatus) {
// 	g.VideoStatusMap[videoStatus.Id] = &videoStatus
// }

// func (g *GlobalVideoStatus) UpdateVideoStatus(videoStatus types.VideoStatus) {
// 	// only update status, percent and alreadyDownloadedSize
// 	g.VideoStatusMap[videoStatus.Id].Title = videoStatus.Title
// 	g.VideoStatusMap[videoStatus.Id].Percent = videoStatus.Percent
// 	g.VideoStatusMap[videoStatus.Id].Status = videoStatus.Status
// 	g.VideoStatusMap[videoStatus.Id].AlreadyDownloadedSize = videoStatus.AlreadyDownloadedSize
// }

// func SaveGlobalVideoStatusMap() {
// 	cwd, err := os.Getwd()
// 	// handle err
// 	if err != nil {
// 		panic(err)
// 	}

// 	raw, err := yaml.Marshal(GlobalVideoStatusMap)
// 	// handle err

// 	err = os.WriteFile(
// 		path.Join(cwd, "video.yaml"),
// 		raw,
// 		0644,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// handle err

// }

// func LoadGlobalVideoStatusMap() {
// 	cwd, err := os.Getwd()
// 	// handle err
// 	if err != nil {
// 		panic(err)
// 	}

// 	raw, err := os.ReadFile(
// 		path.Join(cwd, "video.yaml"),
// 	)
// 	// handle err

// 	var status GlobalVideoStatus
// 	yaml.Unmarshal(raw, &status)
// 	GlobalVideoStatusMap = &status
// 	if GlobalVideoStatusMap == nil {
// 		GlobalVideoStatusMap = &GlobalVideoStatus{
// 			VideoStatusMap:      make(map[string]*types.VideoStatus),
// 			DownloadingVideoNum: 0,
// 		}
// 	}
// 	if GlobalVideoStatusMap.VideoStatusMap == nil {
// 		GlobalVideoStatusMap.VideoStatusMap = make(map[string]*types.VideoStatus)
// 	}
// }
