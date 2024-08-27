package services

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Name   string
	Artist string
	Album  string
}

func CreateSong() error {
	db, err := gorm.Open(sqlite.Open("testdb"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Create(&Song{
		Name:   "Ginger Claps",
		Artist: "Machine Girl",
		Album:  "wlfgrl",
	})
	return nil
}

func GetSongs() error {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Select(&Song{})

	return nil
}
