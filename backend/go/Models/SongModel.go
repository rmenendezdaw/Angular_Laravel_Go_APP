package Models

type Song struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseDate string `json:"releaseDate"`
	Album       string `json:"album"`
	Duration    string `json:"duration"`
	Genre       string `json:"genre"`
}

func (b *Song) TableName() string {
	return "song"
}
