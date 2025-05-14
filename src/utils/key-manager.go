package utils

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
	"log"
	"os"
	"sync"
	"time"
)

type KeyManager struct {
	mutex           sync.RWMutex
	refreshToken    string
	apiToken        string
	expiry          time.Time
	refreshing      bool
	refreshInterval time.Duration
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

	err := km.refreshKey()
	if err != nil {
		log.Printf("Error initialising key: %v\n", err)
	}

	go km.startAutoRefresh()
	return km
}

func (km *KeyManager) GetKey() (string, error) {
	km.mutex.RLock()
	defer km.mutex.RUnlock()
	if km.apiToken == "" {
		return "", fmt.Errorf("token not initialised")
	}
	return "Bearer " + km.apiToken, nil
}

func (km *KeyManager) startAutoRefresh() {
	ticker := time.NewTicker(km.refreshInterval)
	defer ticker.Stop()

	for range ticker.C {
		if time.Now().After(km.expiry) {
			km.mutex.Lock()
			if !km.refreshing {
				km.refreshing = true
				km.mutex.Unlock()
				err := km.refreshKey()
				km.mutex.Lock()
				km.refreshing = false
				km.mutex.Unlock()
				if err != nil {
					log.Printf("Error refreshing key: %v\n", err)
				}
			} else {
				km.mutex.Unlock()
			}
		} else {
			km.mutex.Unlock()
		}
	}
}

func (km *KeyManager) refreshKey() error {
	newKey, expiry, err := km.fetchNewToken()
	if err != nil {
		return err
	}

	km.mutex.Lock()
	defer km.mutex.Unlock()
	km.apiToken = newKey
	km.expiry = expiry
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
