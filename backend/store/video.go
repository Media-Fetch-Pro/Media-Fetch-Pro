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
	// to check video Id is present
	_, ok := s.VideosInfo[videoInfo.Id]
	if !ok {
		s.VideosInfo[videoInfo.Id] = &videoInfo
	} else {
		// TODO: to do all update work, such as update db and update status by percent

		// only update status of attr of videoInfo
		// s.VideosInfo[videoInfo.Id].Title = videoInfo.Title
		// s.VideosInfo[videoInfo.Id].Status = videoInfo.Status
		// s.VideosInfo[videoInfo.Id].Percent = videoInfo.Percent
		s.VideosInfo[videoInfo.Id] = &videoInfo
	}

	// to check videoInfo is vaild or not
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
