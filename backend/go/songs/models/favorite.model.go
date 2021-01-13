package models

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	"goSongs/common"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"goSongs/controllers"
)

type SongUserModel struct {
	gorm.Model
	UserModel      UserModel
	UserModelID    uint
	FavoriteModels []FavoriteModel `gorm:"ForeignKey:FavoriteByID"`
}

type FavoriteModel struct {
	gorm.Model
	Favorite     Songs
	FavoriteID   uint
	FavoriteBy   SongUserModel
	FavoriteByID uint
}

type UserModel struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
	Type         string  `gorm:"column:type;default:'client'"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&SongUserModel{})
	db.AutoMigrate(&FavoriteModel{})
}

func GetSongUserModel(userModel UserModel) SongUserModel {
	var songUserModel SongUserModel
	if userModel.ID == 0 {
		return songUserModel
	}
	db := common.GetDB()
	db.Where(&SongUserModel{
		UserModelID: userModel.ID,
	}).FirstOrCreate(&songUserModel)
	songUserModel.UserModel = userModel
	return songUserModel
}

func (song Songs) isFavoriteBy(user SongUserModel) bool {
	db := common.GetDB()
	var favorite FavoriteModel
	db.Where(FavoriteModel{
		FavoriteID:   song.Id,
		FavoriteByID: user.ID,
	}).First(&favorite)
	return favorite.ID != 0
}

func (song Songs) favoriteBy(user SongUserModel) error {
	db := common.GetDB()
	var favorite FavoriteModel
	err := db.FirstOrCreate(&favorite, &FavoriteModel{
		FavoriteID:   song.Id,
		FavoriteByID: user.ID,
	}).Error
	return err
}

func (song Songs) unFavoriteBy(user SongUserModel) error {
	db := common.GetDB()
	err := db.Unscoped().Where(FavoriteModel{
		FavoriteID:   song.Id,
		FavoriteByID: user.ID,
	}).Delete(FavoriteModel{}).Error
	return err
}

func SongFavorite(c *gin.Context) {
	id := c.Params.ByName("id")
	var song Songs
	err := controllers.GetSongByID(&song, id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("songs", errors.New("Invalid ID")))
		return
	}
	myUserModel := c.MustGet("my_user_model").(UserModel)
	err = song.favoriteBy(GetSongUserModel(myUserModel))
	// serializer := ArticleSerializer{c, articleModel}
	// c.JSON(http.StatusOK, gin.H{"song": serializer.Response()})
	c.JSON(http.StatusOK, gin.H{"song": "test"})

}

func SongUnfavorite(c *gin.Context) {
	id := c.Params.ByName("id")
	var song Songs
	err := controllers.GetSongByID(&song, id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("songs", errors.New("Invalid ID")))
		return
	}
	myUserModel := c.MustGet("my_user_model").(UserModel)
	err = song.unFavoriteBy(GetSongUserModel(myUserModel))
	// serializer := ArticleSerializer{c, articleModel}
	c.JSON(http.StatusOK, gin.H{"song": "delete"})
}