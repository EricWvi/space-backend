package model

import (
	"errors"
	"gorm.io/gorm"
)

const (
	MalFileLocation = iota
	LocalOSS
)

type File struct {
	gorm.Model
	FileField
}

const (
	File_Table    = "files"
	File_Sid      = "sid"
	File_Name     = "name"
	File_Size     = "size"
	File_Link     = "link"
	File_Location = "location"
	File_Type     = "type"
)

func ParseFileLocation(t string) int {
	switch t {
	case "localoss":
		return LocalOSS
	}
	return MalFileLocation
}

func FormatFileLocation(location int) string {
	switch location {
	case LocalOSS:
		return "localoss"
	}
	return ""
}

type FileField struct {
	Sid      Sid    `json:"sid"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Link     string `json:"link"`
	Location int    `json:"location"`
	Type     string `json:"type"`
}

func (f *File) Create(db *gorm.DB) error {
	return db.Create(f).Error
}

func GetFile(db *gorm.DB, where map[string]any) (f *File, err error) {
	var files []File
	db.Table(File_Table).Where(where).Find(&files)
	if len(files) != 1 {
		err = errors.New("file not found")
		return
	}
	f = &files[0]
	return
}
