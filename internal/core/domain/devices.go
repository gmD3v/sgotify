package domain

type Device struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"is_active"`
	Volume int    `json:"volume_percent"`
}
