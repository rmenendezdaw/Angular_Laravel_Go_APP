package routers

import (
	"goUsers/models"
	

	"github.com/gin-gonic/gin"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", models.UsersRegistration)
	router.POST("/login", models.UsersLogin)
}

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", models.UserRetrieve)
	router.PUT("/", models.UserUpdate)
}

func ProfileRegister(router *gin.RouterGroup) {
	router.GET("/:username", models.ProfileRetrieve)
	router.POST("/:username/follow", models.ProfileFollow)
	router.DELETE("/:username/follow", models.ProfileUnfollow)
}

