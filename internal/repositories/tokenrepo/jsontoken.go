package tokenrepo

import (
	"encoding/json"
	"os"
)

type JsonTokenRepository struct{}

func NewTokenRepository() *JsonTokenRepository {
	return &JsonTokenRepository{}
}

type TokenFile struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// SaveToken is a function that saves a token
func (w *JsonTokenRepository) SaveToken(accessToken, refreshToken string) error {
	_, err := os.Stat("token.json")
	if os.IsNotExist(err) {
		err = WriteTokenToFile(accessToken, refreshToken)
	} else {
		err = os.Remove("token.json")
		if err != nil {
			return err
		}
		err = WriteTokenToFile(accessToken, refreshToken)
	}
	if err != nil {
		return err
	}
	return nil
}

func WriteTokenToFile(accessToken, refreshToken string) error {
	file, err := os.Create("token.json")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(`{"access_token":"` + accessToken + `","refresh_token":"` + refreshToken + `"}`)
	if err != nil {
		return err
	}
	return nil
}

// GetSavedToken is a function that returns a token and a refresh token
func (w *JsonTokenRepository) GetSavedToken() (string, string, error) {
	file, err := os.ReadFile("token.json")
	if err != nil {
		return "", "", err
	}
	var token TokenFile
	err = json.Unmarshal(file, &token)

	if err != nil {
		return "", "", err
	}
	return token.AccessToken, token.RefreshToken, nil
}

// DeleteToken is a function that deletes a token
func (w *JsonTokenRepository) DeleteToken() error {
	err := os.Remove("token.json")
	if err != nil {
		return err
	}
	return nil
}
