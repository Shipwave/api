package services

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Id     uint8
	Title  string
	Artist string
	Album  string
}

func CreateSong() error {
	db, err := gorm.Open(sqlite.Open("shipwave.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Create(&Song{
		Id:     1,
		Title:  "Ginger Claps",
		Artist: "Machine Girl",
		Album:  "wlfgrl",
	})
	return nil
}

func GetSongs() error {
	db, err := gorm.Open(sqlite.Open("shipwave.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Select(&Song{Id: 1})

	return nil
}
