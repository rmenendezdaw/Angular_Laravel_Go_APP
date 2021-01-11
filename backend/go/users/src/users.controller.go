package src

import (
	"goUsers/common"
	// "errors"
	// "os"
)

func FindOneUser(condition interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel
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
// Function for validate if exist the same email or username
// func UserRegisterCheck(user *models.UserModel, username string, email string) (error) {
// 	err := Config.DB.Where("username = ?", username).Or("email = ?", email).First(user).Error
// 	if err != nil{
// 		return err
// 	}
// 	return nil
// }
// You could update properties of an UserModel to database returning with error info.
//  err := db.Model(userModel).Update(UserModel{Username: "wangzitian0"}).Error
