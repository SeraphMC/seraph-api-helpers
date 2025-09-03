package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/carlmjohnson/requests"
)

type KeyManager struct {
	mutex           sync.RWMutex
	refreshToken    string
	apiToken        string
	expiry          time.Time
	refreshing      bool
	refreshInterval time.Duration
	errorMsg        error
}

type KeyResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

func NewKeyManager(refreshInterval time.Duration) *KeyManager {
	km := &KeyManager{
		refreshInterval: refreshInterval,
		refreshToken:    os.Getenv("SERAPH_REFRESH_TOKEN"),
	}

	err := km.RefreshToken()
	if err != nil {
		log.Printf("Error initialising key: %v\n", err)
		km.errorMsg = err
	}

	go km.startAutoRefresh()
	return km
}

func (km *KeyManager) GetKey() (string, error) {
	km.tryRefreshKey()
	if km.apiToken == "" {
		return "", fmt.Errorf("token not initialised")
	}
	return "Bearer " + km.apiToken, nil
}

func (km *KeyManager) startAutoRefresh() {
	ticker := time.NewTicker(km.refreshInterval)
	defer ticker.Stop()
	for range ticker.C {
		km.tryRefreshKey()
	}
}

func (km *KeyManager) tryRefreshKey() {
	km.mutex.RLock()
	expired := time.Now().After(km.expiry)
	km.mutex.RUnlock()

	if expired {
		km.mutex.Lock()
		if time.Now().After(km.expiry) && !km.refreshing {
			km.refreshing = true
			defer func() {
				km.refreshing = false
				km.mutex.Unlock()
			}()
			err := km.RefreshToken()
			if err != nil {
				log.Printf("Error refreshing key: %v\n", err)
				km.errorMsg = err
			}
		} else {
			km.mutex.Unlock()
		}
	}
}

func (km *KeyManager) RefreshToken() error {
	newKey, expiry, err := km.fetchNewToken()
	if err != nil {
		return err
	}

	km.mutex.Lock()
	km.apiToken = newKey
	km.expiry = expiry
	km.mutex.Unlock()
	return nil
}

func (km *KeyManager) fetchNewToken() (string, time.Time, error) {
	requestBody := map[string]string{"token": km.refreshToken}
	keyResponse := new(KeyResponse)

	err := requests.URL("https://stash.seraph.si").Path("/auth/refresh").BodyJSON(requestBody).ToJSON(keyResponse).Fetch(context.Background())
	if err != nil {
		return "", time.Time{}, err
	}

	expiresAt, err := time.Parse(time.RFC3339, keyResponse.ExpiresAt)
	if err != nil {
		return "", time.Time{}, err
	}

	return keyResponse.Token, expiresAt, nil
}

func (km *KeyManager) GetError() error {
	return km.errorMsg
}
