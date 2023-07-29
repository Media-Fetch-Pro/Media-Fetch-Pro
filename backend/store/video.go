package store

import (
	"errors"

	"github.com/CorrectRoadH/video-tools-for-nas/backend/types"
)

func (s *Store) AddVideoToQueue(videoInfo types.VideoInfo) {
	s.VideosInfo[videoInfo.Id] = &videoInfo
}

func (s *Store) GetVideoInfoFromQueue(videoId string) *types.VideoInfo {
	return s.VideosInfo[videoId]
}

func (s *Store) UpdateVideoInfo(videoInfo types.VideoInfo) error {
	s.VideosInfo[videoInfo.Id] = &videoInfo
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
