package domain

type Song struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	ReleaseDate string `json:"release_date"`
	Length      int    `json:"length"`
	URI         string `json:"uri"`
	IsPlaying   bool   `json:"is_playing"`
}
