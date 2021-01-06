package models

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
)

type SongsUserModel struct {
	gorm.Model
	UserModel      UserModel
	UserModelID    uint
	FavoriteModels []FavoriteModel `gorm:"ForeignKey:FavoriteByID"`
}

type FavoriteModel struct {
	gorm.Model
	Favorite     Songs
	FavoriteID   uint
	FavoriteBy   SongsUserModel
	FavoriteByID uint
}