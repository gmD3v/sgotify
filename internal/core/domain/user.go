package domain

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
	Expire  int    `json:"expire"`
}
