package caching

import (
	"github.com/Clemintina/common_utilities-for-apis/src/validation"
	"github.com/goccy/go-json"
	"github.com/gofiber/storage/redis"
	"strings"
	"time"
)

func GetFromCache[T interface{}](store *redis.Storage, cacheName string, customStrut T, playerUuid string, callback func(store *redis.Storage) *T) *T {
	redisCache, err := store.Get(strings.ToLower(cacheName + ":" + validation.FormatString(playerUuid)))

	if cacheName == "test" {
		return callback(store)
	}

	if redisCache != nil && err == nil {
		commonCache := &customStrut
		err = json.Unmarshal(redisCache, &commonCache)
		return &customStrut
	}

	return callback(store)
}

func AddToCacheOptional(store *redis.Storage, cacheName string, playerUuid string, customStrut interface{}, cacheTime time.Duration) error {
	if err := AddToRedisCache(store, cacheName, playerUuid, customStrut, cacheTime); err != nil {
		return err
	}
	return nil
}

func AddToRedisCache(store *redis.Storage, cacheName string, playerUuid string, customStrut interface{}, cacheTime time.Duration) error {
	player, _ := json.Marshal(&customStrut)
	if err := store.Set(strings.ToLower(cacheName+":"+validation.FormatString(playerUuid)), player, cacheTime); err != nil {
		return err
	}
	return nil
}
