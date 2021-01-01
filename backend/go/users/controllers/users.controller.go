package controllers

import (
	"goUsers/common"
	"goUsers/models"
	// "errors"
	// "os"
)

func FindOneUser(condition interface{}) (models.UserModel, error) {
	db := common.GetDB()
	var model models.UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// You could input an UserModel which will be saved in database returning with error info
// 	if err := SaveOne(&userModel); err != nil { ... }
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

// You could update properties of an UserModel to database returning with error info.
//  err := db.Model(userModel).Update(UserModel{Username: "wangzitian0"}).Error
