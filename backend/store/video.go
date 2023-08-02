package store

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/types"
)

func (s *Store) AddVideoToQueue(videoInfo types.VideoInfo) {
	s.SchedulerLock.Lock()
	s.VideosInfo[videoInfo.Id] = &videoInfo
	s.SchedulerLock.Unlock()
}

func (s *Store) GetVideoInfoFromQueue(videoId string) *types.VideoInfo {
	return s.VideosInfo[videoId]
}

func (s *Store) UpdateVideoInfo(videoInfo types.VideoInfo) error {
	// s.SchedulerLock.Lock()
	s.VideosInfo[videoInfo.Id] = &videoInfo
	// s.SchedulerLock.Unlock()
	return nil
}

func (s *Store) UpdateVideoInfoPartition(videoInfo types.VideoInfo) error {
	if _, ok := s.VideosInfo[videoInfo.Id]; !ok {
		fmt.Println("new video info:", videoInfo)
		s.VideosInfo[videoInfo.Id] = &videoInfo
	} else {
		willUpdateValue := reflect.ValueOf(videoInfo).Elem()
		beUpdateValue := reflect.ValueOf(s.VideosInfo[videoInfo.Id]).Elem()

		for i := 0; i < willUpdateValue.NumField(); i++ {
			destField := beUpdateValue.Field(i)
			srcField := willUpdateValue.Field(i)

			destField.Set(srcField)
		}
		fmt.Println("new video info:", beUpdateValue)
	}
	s.UpdateChannel <- *s.VideosInfo[videoInfo.Id]
	return nil
}

func (s *Store) GetAllVideoInfo() map[string]*types.VideoInfo {
	return s.VideosInfo
}

func (s *Store) GetVideoInfo(videoId string) (types.VideoInfo, error) {
	result, ok := s.VideosInfo[videoId]
	if ok {
		return *result, nil
	} else {
		return types.VideoInfo{}, errors.New("video not found")
	}
}
