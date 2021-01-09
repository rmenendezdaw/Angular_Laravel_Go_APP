package controllers

import (
	"goSongs/common"
	"github.com/jinzhu/gorm"
	// "errors"
	"fmt"
	// "os"
)

//Create a Song
func CreateSong(data interface{}) error {
	db := common.GetDB()
	err := db.Create(data).Error
	return err
}

//Get all
func GetAllSongs(data interface{}) error {
	fmt.Println(data)
	db := common.GetDB()
	err := db.Find(data).Error
	return err
}

//Update song
func UpdateSong(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

//Delete song
func DeleteSong(data, id interface{}) error {
	db := common.GetDB()
	err := db.Where("id = ?", id).Delete(data).Error
	return err
}

// Delete all songs
func DeleteAllSongs(data interface{}) error {
	db := common.GetDB()
	err := db.Exec(`TRUNCATE TABLE songs`).Error
	return err
}

//Get song by ID
func GetSongByID(data, id interface{}) (error) {
	db := common.GetDB()
	err := db.Where("id = ?", id).First(data).Error
	return err
}
//Increment views of Songs
func IncSongViews(id interface{}) error {
    db := common.GetDB()
    err := db.Table("songs").Where("id = ?", id).Update("views", gorm.Expr("views + 1")).Error
    return err
}
