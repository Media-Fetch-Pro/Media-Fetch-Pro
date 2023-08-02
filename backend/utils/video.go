package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/types"
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

// only run fetch script. the video info will update in the request by script.
func FetchingVideoInfo(videoInfo *types.VideoInfo) error {
	args := []string{"main.py", "--url", videoInfo.Url, "--type", "fetchVideoInfo", "--storage", "/Users/ctrdh/Video"}
	// see the command
	// python3 main.py --url https://www.bilibili.com/video/BV1SW4y1v7WN --type fetchVideoInfo --storage /Users/ctrdh/Video --website bilibili
	out, err := exec.Command("python3", args...).Output()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Printf("out: %s\n", out)
	}
	return err
}

func DownloadVideo(videoInfo *types.VideoInfo, storagePath string) error {
	videoJson, err := json.Marshal(videoInfo)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	args := []string{"main.py", "--url", videoInfo.Url, "--type", "downloadVideo", "--video-info", string(videoJson), "--storage", storagePath}
	out, err := exec.Command("python3", args...).Output()
	if err != nil {
		fmt.Printf("json: %s\n", string(videoJson))
		fmt.Printf("err: %v\n", err)
		fmt.Printf("out: %s\n", out)
	}
	return err
}

func RenameVideo(videoInfo *types.VideoInfo, storagePath string) error {
	videoJson, err := json.Marshal(videoInfo)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	args := []string{"main.py", "--url", videoInfo.Url, "--type", "rename", "--video-info", string(videoJson), "--storage", storagePath}
	out, err := exec.Command("python3", args...).Output()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Printf("out: %s\n", out)
	}
	return err
}
