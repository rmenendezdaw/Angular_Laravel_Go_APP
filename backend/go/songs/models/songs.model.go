package songs

type Songs struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseDate string `json:"release_date"`
	Album       string `json:"album"`
	Duration    string `json:"duration"`
	Genre       string `json:"genre"`
}

