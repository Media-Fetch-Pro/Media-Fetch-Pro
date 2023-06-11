package store

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type SystemSetting struct {
	StoragePath    string
	maxDownloadNum int
	maxSpeed       int
}

var SystemSettingCache = &SystemSetting{
	StoragePath:    "/home/ctrdh/video",
	maxDownloadNum: 5,
	maxSpeed:       100,
}

func LoadSystemSetting() SystemSetting {
	cwd, err := os.Getwd()
	// handle err
	if err != nil {
		panic(err)
	}

	raw, err := os.ReadFile(
		path.Join(cwd, "config.yaml"),
	)
	// handle err

	var config SystemSetting
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
