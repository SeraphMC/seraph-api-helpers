package caching

import (
	"bytes"
	"compress/zlib"
	"github.com/SeraphMC/seraph-api-helpers/src/validation"
	"github.com/goccy/go-json"
	"github.com/gofiber/storage/redis"
	"strings"
	"time"
)

func Compress(data []byte) ([]byte, error) {
	byteBuffer := new(bytes.Buffer)
	writer, _ := zlib.NewWriterLevel(byteBuffer, 6)

	_, err := writer.Write(data)
	if err != nil {
		_ = writer.Close()
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return byteBuffer.Bytes(), nil
}

func Decompress(data []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	out := new(bytes.Buffer)
	_, err = out.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func GetFromCache[T any](store *redis.Storage, cacheName string, customStruct T, playerUuid string, callback func(store *redis.Storage) *T) *T {
	redisCache, err := store.Get(strings.ToLower(cacheName + ":" + validation.FormatString(playerUuid)))

	if cacheName == "test" {
		return callback(store)
	}

	if redisCache != nil && err == nil {
		decompressedCache, err := Decompress(redisCache)
		if err != nil {
			return callback(store)
		}

		commonCache := &customStruct
		err = json.Unmarshal(decompressedCache, &commonCache)
		if err == nil {
			return &customStruct
		}
	}

	return callback(store)
}

func AddToCacheOptional(store *redis.Storage, cacheName string, playerUuid string, customStruct interface{}, cacheTime time.Duration) error {
	if err := AddToRedisCache(store, cacheName, playerUuid, customStruct, cacheTime); err != nil {
		return err
	}
	return nil
}

func AddToRedisCache(store *redis.Storage, cacheName string, playerUuid string, customStruct interface{}, cacheTime time.Duration) error {
	playerData, err := json.Marshal(&customStruct)
	if err != nil {
		return err
	}

	compressedData, err := Compress(playerData)
	if err != nil {
		return err
	}

	if err := store.Set(strings.ToLower(cacheName+":"+validation.FormatString(playerUuid)), compressedData, cacheTime); err != nil {
		return err
	}
	return nil
}
