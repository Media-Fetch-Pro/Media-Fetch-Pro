package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/CorrectRoadH/video-tools-for-nas/backend/types"
)

func HandlerDownloader(url string) (string, error) {
	// 判断url是bilibili还是youtube
	// 如果是bilibili，则返回bilibili
	// 如果是youtube，则返回youtube
	// 如果是其他，返回错误
	// 这里用正则做判断，如 https://www.youtube.com/watch?v=KO-yFbjtXGg
	// https://www.bilibili.com/video/BV1Y7411t7zZ
	youtubeRegex := regexp.MustCompile(`https?://www.youtube.com/watch\?v=.*`)
	bilibiliRegex := regexp.MustCompile(`https?://www.bilibili.com/video/.*`)

	if youtubeRegex.MatchString(url) {
		return "youtube", nil
	} else if bilibiliRegex.MatchString(url) {
		return "bilibili", nil
	} else {
		return "unknown", fmt.Errorf("unknown website")
	}
}

func GenerateVideoIdFromURL(url string) string {
	// return md5 of url
	data := []byte(url)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func FetchingVideoInfo(videoInfo *types.VideoInfo) (*types.VideoInfo, error) {
	fmt.Printf("start fetching videoInfo: %v\n", videoInfo)
	args := []string{"main.py", "--url", videoInfo.Url, "--type", "fetchVideoInfo", "--storage", "/Users/ctrdh/Video", "--website", "bilibili"}
	out, err := exec.Command("python3", args...).Output()

	// convert json of out to videoInfo
	var videoInfos []types.VideoInfo
	json.Unmarshal(out, &videoInfos)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Printf("out: %s\n", out)
	}

	fmt.Println("out:")
	fmt.Println(videoInfos)

	// for _, v := range videoInfos {
	// 	server.Store.UpdateVideoInfo(v)
	// }

	return videoInfo, err
}

func DownloadVideo(videoInfo *types.VideoInfo, storagePath string) error {
	fmt.Printf("videoInfo: %v\n", videoInfo)
	args := []string{"main.py", "--url", videoInfo.Url, "--storage", storagePath, "--type", videoInfo.Type}
	out, err := exec.Command("python3", args...).Output()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Printf("out: %s\n", out)
	}
	return err
}
