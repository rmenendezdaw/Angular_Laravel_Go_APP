package songs

import (
	"fmt"
	"app/common"
	"strconv"
	"github.com/jinzhu/gorm"
)

type Songs struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseDate string `json:"releaseDate"`
	Album       string `json:"album"`
	Duration    string `json:"duration"`
	Genre       string `json:"genre"`
}
}
