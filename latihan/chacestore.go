package latihan

import (
	"errors"
	"fmt"
)

type CacheStore struct {
	Data map[string]string
}

func NewCacheStore() *CacheStore {
	return &CacheStore{
		make(map[string]string),
	}
}

func (c *CacheStore) GetUserName(userID string) (string, error) {
	if value, ok := c.Data[userID]; ok {
		return fmt.Sprintf("%s dari cache \n", value), nil
	}
	return "", errors.New("user tidak ditemukan di dalam cache")
}
