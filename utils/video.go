package utils

import (
	"fmt"
	"regexp"
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
