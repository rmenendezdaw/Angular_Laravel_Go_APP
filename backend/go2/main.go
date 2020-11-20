package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rmenendezdaw/Angular_Laravel_Go_APP/backend/go2/songs"
	"github.com/rmenendezdaw/Angular_Laravel_Go_APP/backend/go2/articles"
	"github.com/rmenendezdaw/Angular_Laravel_Go_APP/backend/go2/common"
	"github.com/rmenendezdaw/Angular_Laravel_Go_APP/backend/go2/users"
)

func Migrate(db *gorm.DB) {
	// users.AutoMigrate()
	// db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&songs.Songs{})
	// db.AutoMigrate(&articles.TagModel{})
	// db.AutoMigrate(&articles.FavoriteModel{})
	// db.AutoMigrate(&articles.ArticleUserModel{})
	// db.AutoMigrate(&articles.CommentModel{})
}

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	MakeRoutes(r)

	v1 := r.Group("/api")

	songs.SongsRegister(v1.Group("/songs"))
	songs.SongsAnonymousRegister(v1.Group("/songs"))

	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))
	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))
	articles.ArticlesRegister(v1.Group("/articles"))

	fmt.Printf("0.0.0.0:3000")
	r.Run(":3000")
}

func MakeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()

		
			fmt.Printf("c.Request.RequestURI \n")
			fmt.Printf(c.Request.RequestURI)
	}
	r.Use(cors)
}
