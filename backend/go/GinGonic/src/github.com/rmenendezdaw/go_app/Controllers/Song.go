package Controllers

import (
	"fmt"
	"go_app/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetSongs ... Get all songs
func GetSongs(c *gin.Context) {
	var song []Models.Song
	err := Models.GetAllSongs(&song)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, song)
	}
}

//CreateSong ... Create Song
func CreateSong(c *gin.Context) {
	var song Models.Song
	c.BindJSON(&song)
	err := Models.CreateSong(&song)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, song)
	}
}

//GetSongByID ... Get the song by id
func GetSongByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var song Models.Song
	err := Models.GetSongByID(&song, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, song)
	}
}

//UpdateSong ... Update the song information
func UpdateSong(c *gin.Context) {
	var song Models.Song
	id := c.Params.ByName("id")
	err := Models.GetSongByID(&song, id)
	if err != nil {
		c.JSON(http.StatusNotFound, song)
	}
	c.BindJSON(&song)
	err = Models.UpdateSong(&song, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, song)
	}
}

//DeleteSong ... Delete the song
func DeleteSong(c *gin.Context) {
	var song Models.Song
	id := c.Params.ByName("id")
	err := Models.DeleteSong(&song, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
