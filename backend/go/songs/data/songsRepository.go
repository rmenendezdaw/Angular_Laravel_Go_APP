func SongCreate(c *gin.Context) {
	var song Songs
	c.BindJSON(&song);

	err := CreateSong(&song)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"song": song})
		return
	}
}

func SongList(c *gin.Context) {
	var song []Songs
	err := GetAllSongs(&song)
	if err != nil {
		c.JSON(http.StatusOK, "Not found")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, song)
	}
}

func SongById(c *gin.Context) {
	id := c.Params.ByName("id")
	var song Songs
	err := GetSongByID(&song, id)
	if err != nil {
		c.JSON(http.StatusOK, "Not found")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"song": song})
		return
	}
}

func SongUpdate(c *gin.Context) {
	var song Songs
	id := c.Params.ByName("id")
	err := GetSongByID(&song, id) 
	if err != nil { 
		c.JSON(http.StatusNotFound, "NOT FOUND")
	}else{ 
		c.BindJSON(&song)
		err = UpdateSong(&song) 
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
	var song Songs
	id := c.Params.ByName("id")
	err := DeleteSong(&song, id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("songs", errors.New("Invalid slug")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"song": "Delete success"})
}
func SongDeleteAll(c *gin.Context) {
	var song Songs
	err := DeleteAllSongs(&song)
	if err != nil {
		c.JSON(http.StatusOK, "Not found")
	} else {
		c.JSON(http.StatusOK, "DELETED ALL songs")
	}
}