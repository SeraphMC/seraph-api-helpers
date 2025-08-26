package minecraft

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

// MojangResponse represents the response structure received from Mojang API for user-related queries.
// It includes the user's unique ID, username, an error message if applicable, and the request path.
type MojangResponse struct {
	Id           string `json:"id,omitempty"`
	Username     string `json:"name,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Path         string `json:"path,omitempty"`
}

// FetchMojangProfile retrieves a Mojang profile based on the provided Minecraft username.
// It returns a pointer to a MojangResponse struct containing profile details or nil if the request fails.
func FetchMojangProfile(username string) *MojangResponse {
	playerStruct := new(MojangResponse)
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", username)
	err := requests.URL(url).ToJSON(playerStruct).Fetch(context.Background())
	if err != nil {
		return nil
	}
	return playerStruct
}

// FetchMojangSessionProfile retrieves a Mojang player's profile data using their UUID from Mojang's session server.
// It returns a pointer to a MojangResponse struct or nil if an error occurs during the request.
func FetchMojangSessionProfile(uuid string) *MojangResponse {
	playerStruct := new(MojangResponse)
	url := fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/profile/%s", uuid)
	err := requests.URL(url).ToJSON(playerStruct).Fetch(context.Background())
	if err != nil {
		return nil
	}
	return playerStruct
}
