package store_test

import (
	"fmt"
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
	})
	fmt.Println("url:", s.GetVideoInfoFromQueue("test").Url)

	s.UpdateVideoInfoPartition(types.VideoInfo{
		Id:     "test",
		Status: "test2",
	})

	assert.Equal(t, s.GetVideoInfoFromQueue("test").Status, "test2")
	fmt.Println("url:", s.GetVideoInfoFromQueue("test").Url)
	assert.Equal(t, s.GetVideoInfoFromQueue("test").Url, "test")
	assert.Equal(t, s.GetVideoInfoFromQueue("test").Type, "test")
	assert.Equal(t, s.GetVideoInfoFromQueue("test").Length, nil)

	s.UpdateVideoInfoPartition(types.VideoInfo{
		Id:     "test2",
		Status: "test2",
	})

	assert.Equal(t, s.GetVideoInfoFromQueue("test2").Status, "test2")
	assert.Equal(t, s.GetVideoInfoFromQueue("test2").Status, "test2")
	assert.Equal(t, s.GetVideoInfoFromQueue("test2").Url, nil)

}
