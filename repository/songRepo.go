package repository

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Id     uint
	Title  string
	Artist string
	Album  string
}

type SongRepo struct {
	db *gorm.DB
}

func NewSongRepo(db *gorm.DB) *SongRepo {
	return &SongRepo{db: db}
}
