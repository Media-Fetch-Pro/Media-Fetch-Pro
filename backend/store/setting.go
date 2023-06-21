package store

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type SystemSettings struct {
	StoragePath    string
	MaxDownloadNum int
	MaxSpeed       int
}

var SystemSettingCache = &SystemSettings{
	StoragePath:    "/var/opt/video",
	MaxDownloadNum: 5,
	MaxSpeed:       100,
}

func LoadSystemSetting() SystemSettings {
	cwd, err := os.Getwd()
	// handle err
	if err != nil {
		panic(err)
	}

	raw, err := os.ReadFile(
		path.Join(cwd, "config.yaml"),
	)
	// handle err

	var config SystemSettings
	yaml.Unmarshal(raw, &config)

	return config
}

func SaveSystemSetting() {
	cwd, err := os.Getwd()
	// handle err
	if err != nil {
		panic(err)
	}

	raw, err := yaml.Marshal(SystemSettingCache)
	// handle err

	err = os.WriteFile(
		path.Join(cwd, "config.yaml"),
		raw,
		0644,
	)
	if err != nil {
		panic(err)
	}
	// handle err
}
