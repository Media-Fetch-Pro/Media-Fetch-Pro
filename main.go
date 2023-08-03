package main

import (
	"fmt"

	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/server"
	_profile "github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/server/profile"
	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/store"
	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/store/db"
	"github.com/spf13/viper"
)

const (
	greetingBanner = `
	+-+-+-+-+-+ +-+-+-+-+-+
	|M|e|d|i|a| |F|e|t|c|h|
	+-+-+-+-+-+ +-+-+-+-+-+
`
)

var profile *_profile.Profile

func initConfig() {
	viper.AutomaticEnv()
	var err error
	profile, err = _profile.GetProfile()
	if err != nil {
		fmt.Printf("failed to get profile, error: %+v\n", err)
		return
	}

	println("---")
	println("Server profile")
	println("dsn:", profile.DSN)
	println("port:", profile.Port)
	println("mode:", profile.Mode)
	println("version:", profile.Version)
	println("---")
}

func main() {
	initConfig()

	db := db.NewDB(profile)
	storeInstance := store.NewStore(db.DBInstance)

	s, err := server.NewServer(profile, storeInstance)
	if err != nil {
		panic(err)
	}

	println(greetingBanner)

	s.Run(":7789")
}
