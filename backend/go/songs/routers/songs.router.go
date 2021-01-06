package routers

import (
	"goSongs/data"
	"github.com/gin-gonic/gin"
)

func SongsRegister(router *gin.RouterGroup) {
	router.POST("/", data.SongCreate)
	router.PUT("/:id", data.SongUpdate)
	router.DELETE("/:id", data.SongDelete)
	router.DELETE("/", data.SongDeleteAll)
	router.POST("/:id/favorite", SongFavorite)
	router.DELETE("/:id/favorite", SongUnfavorite)

}

func SongsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", data.SongList)
	router.GET("/:id", data.SongById)
}

