package main

import (
	"5_go_mysql/Config"
	"5_go_mysql/Models"
	"5_go_mysql/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()
	//running
	// Listen and Server in 0.0.0.0:8080
	r.Run(":3000")
}
