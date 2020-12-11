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

}

func SongsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", data.SongList)
	router.GET("/:id", data.SongById)
}