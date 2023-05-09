package buz

import (
	"github.com/izhujiang/appengine/module/core"
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  core.UUID
}
