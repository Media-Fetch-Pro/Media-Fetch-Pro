package store

import (
	"errors"
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

// if to set the value to default. It will didn't have effect.
// But I think it is not a problem. Because we seem didn't have the case to set the value to default.
func isFieldEmpty(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.String:
		return field.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() == 0
	case reflect.Array, reflect.Slice:
		return field.Len() == 0
	default:
		return false // Add more cases as needed for other data types
	}
}

func (s *Store) UpdateVideoInfoPartition(videoInfo types.VideoInfo) error {
	if _, ok := s.VideosInfo[videoInfo.Id]; !ok {
		s.VideosInfo[videoInfo.Id] = &videoInfo
	} else {
		dest := videoInfo
		src := *(s.VideosInfo[videoInfo.Id])
		destValue := reflect.ValueOf(&dest).Elem()
		srcValue := reflect.ValueOf(&src).Elem()

		for i := 0; i < destValue.NumField(); i++ {
			destField := destValue.Field(i)
			srcField := srcValue.Field(i)

			if !isFieldEmpty(destField) {
				srcField.Set(destField)
			}
		}
		s.VideosInfo[videoInfo.Id] = &src
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
