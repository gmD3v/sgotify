package domain

type Playlist struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Owner           string `json:"owner"`
	URI             string `json:"uri"`
	IsPublic        bool   `json:"is_public"`
	IsCollaborative bool   `json:"is_collaborative"`
	Tracks          []Song `json:"tracks"`
}
