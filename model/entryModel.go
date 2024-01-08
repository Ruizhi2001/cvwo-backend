package model

import (
	"errors"

	"github.com/Ruizhi2001/cvwo-backend/database"
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (entry *Entry) Save() (*Entry, error) {
	if entry.Content == "" {
		return &Entry{}, errors.New("content cannot be empty")
	}

	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
