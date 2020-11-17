package songs

import (
	"errors"
	"net/http"

	"github.com/rmenendezdaw/Angular_Laravel_Go_APP/backend/go2/common"

	"github.com/gin-gonic/gin"
)

func SongsRegister(router *gin.RouterGroup) {
	router.POST("/", SongCreate)
	router.GET("/", SongList)
	router.PUT("/:id", SongUpdate)
	router.DELETE("/:id", SongDelete)
	router.DELETE("/", SongDeleteAll)

}

func SongsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", SongList)
	router.GET("/:id", SongRetrieve)
}

func TagsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", TagList)
}

func SongCreate(c *gin.Context) {
	songModelValidator := NewSongModelValidator()
	if err := songModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	//fmt.Println(songModelValidator.songModel.Author.UserModel)

	if err := SaveOne(&songModelValidator.songModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := SongSerializer{c, songModelValidator.songModel}
	c.JSON(http.StatusCreated, gin.H{"song": serializer.Response()})
}

func SongList(c *gin.Context) {
	//condition := SongModel{}
	tag := c.Query("tag")
	author := c.Query("author")
	favorited := c.Query("favorited")
	limit := c.Query("limit")
	offset := c.Query("offset")
	songModels, modelCount, err := FindManySong(tag, author, limit, offset, favorited)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("songs", errors.New("Invalid param")))
		return
	}
	serializer := SongsSerializer{c, songModels}
	c.JSON(http.StatusOK, gin.H{"songs": serializer.Response(), "songsCount": modelCount})
}

func SongRetrieve(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "feed" {
		SongFeed(c)
		return
	}
	songModel, err := FindOneSong(&SongModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("songs", errors.New("Invalid slug")))
		return
	}
	serializer := SongSerializer{c, songModel}
	c.JSON(http.StatusOK, gin.H{"song": serializer.Response()})
}

func SongUpdate(c *gin.Context) {
	slug := c.Param("slug")
	songModel, err := FindOneSong(&SongModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("songs", errors.New("Invalid slug")))
		return
	}
	songModelValidator := NewSongModelValidatorFillWith(songModel)
	if err := songModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	songModelValidator.songModel.ID = songModel.ID
	if err := songModel.Update(songModelValidator.songModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := SongSerializer{c, songModel}
	c.JSON(http.StatusOK, gin.H{"song": serializer.Response()})
}

func SongDelete(c *gin.Context) {
	slug := c.Param("slug")
	err := DeleteSongModel(&SongModel{Slug: slug})
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
		c.JSON(http.StatusOK, "Truncate song")
	}
}
func TagList(c *gin.Context) {
	tagModels, err := getAllTags()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("songs", errors.New("Invalid param")))
		return
	}
	serializer := TagsSerializer{c, tagModels}
	c.JSON(http.StatusOK, gin.H{"tags": serializer.Response()})
}
