package Routes

import (
	"app/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("song", Controllers.GetSongs)
		grp1.POST("song", Controllers.CreateSong)
		grp1.GET("song/:id", Controllers.GetSongByID)
		grp1.PUT("song/:id", Controllers.UpdateSong)
		grp1.DELETE("song/:id", Controllers.DeleteSong)
	}
	return r
}
