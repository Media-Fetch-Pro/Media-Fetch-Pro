package store_test

import (
	"testing"

	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/store"
	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/types"
	"github.com/go-playground/assert/v2"
)

func TestUpdateVideoInfoPartition(t *testing.T) {
	// new a store
	// new a videoInfo
	// update videoInfo
	s := store.NewStore(nil)
	s.AddVideoToQueue(types.VideoInfo{
		Id:     "test",
		Url:    "test",
		Status: "test",
		Type:   "test",
		Children: []string{
			"1234",
			"234",
		},
	})

	s.UpdateVideoInfoPartition(types.VideoInfo{
		Id:     "test",
		Status: "test2",
	})

	assert.Equal(t, s.GetVideoInfoFromQueue("test").Status, "test2")
	assert.Equal(t, s.GetVideoInfoFromQueue("test").Url, "test")
	assert.Equal(t, s.GetVideoInfoFromQueue("test").Type, "test")
	assert.Equal(t, s.GetVideoInfoFromQueue("test").Length, 0)
	assert.Equal(t, s.GetVideoInfoFromQueue("test").Children, []string{
		"1234",
		"234",
	})

	s.UpdateVideoInfoPartition(types.VideoInfo{
		Id:     "test2",
		Status: "test2",
	})

	assert.Equal(t, s.GetVideoInfoFromQueue("test2").Status, "test2")
	assert.Equal(t, s.GetVideoInfoFromQueue("test2").Status, "test2")
	assert.Equal(t, s.GetVideoInfoFromQueue("test2").Url, "")

	s.AddVideoToQueue(types.VideoInfo{
		Id:     "test3",
		Url:    "test",
		Status: "test",
		Type:   "test",
		Length: 5,
		Children: []string{
			"1234",
			"234",
		},
	})

	// s.UpdateVideoInfoPartition(types.VideoInfo{
	// 	Id:       "test3",
	// 	Children: []string{},
	// 	Url:      "",
	// 	Length:   0,
	// })
	// assert.Equal(t, s.GetVideoInfoFromQueue("test3").Children, []string{})
	// assert.Equal(t, s.GetVideoInfoFromQueue("test3").Url, "")
	// assert.Equal(t, s.GetVideoInfoFromQueue("test3").Length, 0)
}
