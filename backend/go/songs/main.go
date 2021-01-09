package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goSongs/models"
	"goSongs/common"
	"goSongs/routers"


)

func Migrate(db *gorm.DB) {
	models.AutoMigrate()
}

// Entry point for the program
func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	MakeRoutes(r)

	v1 := r.Group("/api")

	v1.Use(models.AuthMiddleware(false))
	routers.SongsAnonymousRegister(v1.Group("/songs"))
	
	v1.Use(models.AuthMiddleware(true))
	routers.SongsRegister(v1.Group("/songs"))

	
	fmt.Printf("0.0.0.0:3001")
	r.Run(":3001")
}

func MakeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		fmt.Printf("c.Request.Method \n")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()

		/*
			fmt.Printf("c.Request.Method \n")
			fmt.Printf(c.Request.Method)
			fmt.Printf("c.Request.RequestURI \n")
			fmt.Printf(c.Request.RequestURI)
		*/
	}
	r.Use(cors)
}