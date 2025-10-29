package caching

import (
	"bytes"
	"compress/zlib"
	"strings"
	"time"

	"github.com/SeraphMC/seraph-api-helpers/src/validation"
	"github.com/goccy/go-json"
	"github.com/gofiber/storage/redis/v3"
)

// Compress compresses the input byte slice using zlib compression with level 6 and returns the compressed data or an error if compression fails.
func Compress(data []byte) ([]byte, error) {
	byteBuffer := new(bytes.Buffer)
	writer, err := zlib.NewWriterLevel(byteBuffer, 6)
	if err != nil {
		return nil, err
	}

	defer writer.Close()

	_, err = writer.Write(data)
	if err != nil {
		return nil, err
	}

	if err = writer.Close(); err != nil {
		return nil, err
	}

	return byteBuffer.Bytes(), nil
}

// Decompress takes a byte slice, decompresses it using zlib, and returns the resulting byte slice or an error if decompression fails.
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

// GetFromCache retrieves a cached value from Redis storage, decompresses it, and unmarshals it into the given structure. If no cache is found or on error, it invokes the provided callback to fetch a new value and returns it.
func GetFromCache[T any](store *redis.Storage, cacheName string, playerUuid string, callback func(store *redis.Storage) *T) *T {
	redisCache, err := store.Get(strings.ToLower(cacheName + ":" + validation.FormatString(playerUuid)))

	if cacheName == "test" {
		return callback(store)
	}

	if redisCache != nil && err == nil {
		decompressedCache, err := Decompress(redisCache)
		if err != nil {
			return callback(store)
		}

		var result T
		err = json.Unmarshal(decompressedCache, &result)
		if err == nil {
			return &result
		}
	}

	return callback(store)
}

// AddToCacheOptional stores a custom data structure in a Redis cache with optional parameters for cache time and compresses the data before storage.
func AddToCacheOptional(store *redis.Storage, cacheName string, playerUuid string, customStruct interface{}, cacheTime time.Duration) error {
	if err := AddToRedisCache(store, cacheName, playerUuid, customStruct, cacheTime); err != nil {
		return err
	}
	return nil
}

// AddToRedisCache stores a compressed, JSON-encoded representation of a customStruct in the Redis cache under a key formed from a combination of cacheName and playerUuid, with a specified expiration time. Returns an error if encoding, compression, or storage fails.
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
