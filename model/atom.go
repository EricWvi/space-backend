package model

import (
	"github.com/space-backend/config"
	"gorm.io/gorm"
)

const (
	MalAtomType = iota
	Text
	Image
	Audio
	Video
)

type Atom struct {
	gorm.Model
	Sid     int64  `json:"sid"`
	Content string `json:"content"`
	Link    string `json:"link"`
	Type    int    `json:"type"`
	DocId   int64  `json:"docId"`
}

func FormatAtomType(t string) int {
	switch t {
	case "text":
		return Text
	case "image":
		return Image
	case "audio":
		return Audio
	case "video":
		return Video
	}
	return MalAtomType
}

func (a *Atom) Create() error {
	return config.DB.Create(a).Error
}
