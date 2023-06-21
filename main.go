package main

import (
	"github.com/CorrectRoadH/video-tools-for-nas/backend/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	s.Run(":7789")
}
