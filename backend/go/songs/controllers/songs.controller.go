package controllers

import (
	"goSongs/common"
	"github.com/jinzhu/gorm"
	
	// "errors"
	"fmt"
	// "os"
	"strconv"
)

//Create a Song
func CreateSong(data interface{}) error {
	db := common.GetDB()
	err := db.Create(data).Error
	return err
}

//Get all
func GetAllSongs(views, release_date, offset, limit string, data interface{}) (error, int) {
	fmt.Println(data)
	var count = 0
	db := common.GetDB()
	db.Model(data).Count(&count)

	offset_int, err := strconv.Atoi(offset)
	if err != nil {
		offset_int = 0
	}

	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		limit_int = count
	}
	

	if (views != "") {
		err = db.Order("views " + views).Offset(offset_int).Limit(limit_int).Find(data).Error
	}else if (release_date != "") {
		err = db.Order("views " + views).Offset(offset_int).Limit(limit_int).Find(data).Error
	} else {
		err = db.Offset(offset_int).Limit(limit_int).Find(data).Error
	}// end_else


	return err, count
	// return err
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
