package minecraft

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
)

type MojangResponse struct {
	Id           string `json:"id,omitempty"`
	Username     string `json:"name,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Path         string `json:"path,omitempty"`
}

func FetchMojangProfile(username string) *MojangResponse {
	playerStruct := new(MojangResponse)
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", username)
	err := requests.URL(url).ToJSON(playerStruct).Fetch(context.Background())
	if err != nil {
		return nil
	}
	return playerStruct
}

func FetchMojangSessionProfile(uuid string) *MojangResponse {
	playerStruct := new(MojangResponse)
	url := fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/profile/%s", uuid)
	err := requests.URL(url).ToJSON(playerStruct).Fetch(context.Background())
	if err != nil {
		return nil
	}
	return playerStruct
}
