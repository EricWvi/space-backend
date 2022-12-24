package service

import (
	"github.com/space-backend/model"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func UploadFile(location int, name string, content []byte) (link string, err error) {
	switch location {
	case model.LocalOSS:
		root := viper.GetString("localoss.location")
		dir1 := name[:2]
		dir2 := name[2:4]
		path := filepath.Join(root, dir1, dir2)
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
		link = filepath.Join(dir1, dir2, name)
		err = os.WriteFile(filepath.Join(root, link), content, 0644)
		if err != nil {
			return
		}
		return
	default:
		return
	}
}

func DownloadFile(location int, link string) (content []byte, err error) {
	switch location {
	case model.LocalOSS:
		root := viper.GetString("localoss.location")
		path := filepath.Join(root, link)
		return os.ReadFile(path)

	default:
		return
	}
}
