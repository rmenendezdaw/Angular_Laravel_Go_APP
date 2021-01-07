package models

import (
	"github.com/gin-gonic/gin"
)

type SongResponse struct {
	Id             uint                  `json:"id"`
	Title          string                `json:"title"`
	Artist          string                `json:"artist"`
	ReleaseDate    string                `json:"release_date"`
	Album           string                `json:"album"`
	Duration      string                `json:"duration"`
	Genre         string 					`json:"genre"`
	Favorite       bool                  `json:"favorited"`
}

type SongsSerializer struct {
	C        *gin.Context
	Songs []Songs
}

type SongSerializer struct {
	C *gin.Context
	Songs
}

func (s *SongSerializer) Response() SongResponse {
	myUserModel := s.C.MustGet("my_user_model").(UserModel)
	response := SongResponse {
		Id:          s.Id,
		Title:       s.Title,
		Artist: 	s.Artist,
		ReleaseDate:        s.ReleaseDate,
		Album:   		s.Album,
		Duration:		s.Duration,
		Genre:         	s.Genre,
		Favorite:       s.isFavoriteBy(GetSongUserModel(myUserModel)),
	}
	
	return response
}

func (s *SongsSerializer) Response() []SongResponse {
	response := []SongResponse{}
	for _, song := range s.Songs {
		serializer := SongSerializer{s.C, song}
		response = append(response, serializer.Response())
	}
	return response
}