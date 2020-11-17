package Models

import (
	"app/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllSongs Fetch all song data
func GetAllSongs(song *[]Song) (err error) {
	if err = Config.DB.Find(song).Error; err != nil {
		return err
	}
	return nil
}

//CreateSong ... Insert New data
func CreateSong(song *Song) (err error) {
	if err = Config.DB.Create(song).Error; err != nil {
		return err
	}
	return nil
}

//GetSongByID ... Fetch only one song by Id
func GetSongByID(song *Song, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(song).Error; err != nil {
		return err
	}
	return nil
}

//UpdateSong ... Update song
func UpdateSong(song *Song, id string) (err error) {
	fmt.Println(song)
	Config.DB.Save(song)
	return nil
}

//DeleteSong ... Delete song
func DeleteSong(song *Song, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(song)
	return nil
}
