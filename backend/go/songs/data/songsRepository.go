package data

import (
	// "time"
	"errors"
	"goSongs/common"
	"goSongs/controllers"
	"goSongs/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SongCreate(c *gin.Context) {
	var song models.Songs
	c.BindJSON(&song)

	err := controllers.CreateSong(&song)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"song": song})
		return
	}
}

func SongList(c *gin.Context) {
	var song []models.Songs
	err := controllers.GetAllSongs(&song)
	
	if err != nil {
		c.JSON(http.StatusOK, "Not found")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		serializer := models.SongsSerializer{c, song}
		c.JSON(http.StatusOK, gin.H{"songs": serializer.Response()})
	}
}

func SongById(c *gin.Context) {
	id := c.Params.ByName("id")
	var song models.Songs
	err := controllers.GetSongByID(&song, id)
	if err != nil {
		c.JSON(http.StatusOK, "Not found")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		serializer := models.SongSerializer{c, song}
		c.JSON(http.StatusOK, gin.H{"song": serializer.Response()})
		return
	}
}

func SongUpdate(c *gin.Context) {
	var song models.Songs
	id := c.Params.ByName("id")
	err := controllers.GetSongByID(&song, id)
	if err != nil {
		c.JSON(http.StatusNotFound, "NOT FOUND")
	} else {
		c.BindJSON(&song)
		err = controllers.UpdateSong(&song)
		if err != nil {
			c.JSON(http.StatusOK, "Not found")
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"song": song})
			return
		}
	}
}

func SongDelete(c *gin.Context) {
	var song models.Songs
	id := c.Params.ByName("id")
	err := controllers.DeleteSong(&song, id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("songs", errors.New("Invalid slug")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"song": "Delete success"})
}
func SongDeleteAll(c *gin.Context) {
	var song models.Songs
	err := controllers.DeleteAllSongs(&song)
	if err != nil {
		c.JSON(http.StatusOK, "Not found")
	} else {
		c.JSON(http.StatusOK, "DELETED ALL songs")
	}
}
