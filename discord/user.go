package discord

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
	Bot           bool   `json:"bot"`
	Banner        string `json:"banner"`
	BannerColor   string `json:"banner_color"`
	AccentColor   string `json:"accent_color"`
}

func GetUser(id string, token string) (User, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://discord.com/api/v9/users/%s", id), nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Add("User-Agent", "Discord Finder")
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	user := User{}
	json.NewDecoder(resp.Body).Decode(&user)
	return user, nil
}
