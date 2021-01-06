package routers

import (
	"goSongs/data"
	"goSongs/models"
	"github.com/gin-gonic/gin"
)

func SongsRegister(router *gin.RouterGroup) {
	router.POST("/", data.SongCreate)
	router.PUT("/:id", data.SongUpdate)
	router.DELETE("/:id", data.SongDelete)
	router.DELETE("/", data.SongDeleteAll)
	router.POST("/:id/favorite", models.SongFavorite)
	// router.DELETE("/:id/favorite", data.SongUnfavorite)

}

func SongsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", data.SongList)
	router.GET("/:id", data.SongById)
}

