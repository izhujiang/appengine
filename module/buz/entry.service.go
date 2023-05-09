package buz

import (
	"github.com/izhujiang/appengine/module/core"
	"github.com/izhujiang/appengine/module/db"
)

func SaveEntry(entry *Entry) (*Entry, error) {
	err := db.Instance.Create(entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

func FindEntries(userID core.UUID) ([]*Entry, error) {
	var entries []*Entry
	err := db.Instance.Debug().Where("user_id = ?", userID).Find(&entries).Error
	if err != nil {
		return entries, err
	}
	return entries, nil
}
