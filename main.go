package main

import (
	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	s.Run(":7789")
}
