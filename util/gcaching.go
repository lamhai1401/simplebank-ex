package util

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

const (
	// ExpireTime linter
	ExpireTime = time.Duration(3600)
)

var localCaching gcache.Cache

// AddLocalCacheErr litner
var AddLocalCacheErr = func(reason string) error {
	return fmt.Errorf("AddLocalCacheErr err: %v", reason)
}

// GetLocalCacheErr linter
var GetLocalCacheErr = func(reason string) error {
	return fmt.Errorf("GetLocalCacheErr err: %v", reason)
}

// InitLocalCaching linter
func InitLocalCaching(size int) {
	localCaching = gcache.New(size).LFU().Build()
}

// AddLocalCache linter
func AddLocalCache(key string, value interface{}) error {
	err := localCaching.SetWithExpire(key, value, ExpireTime*time.Second)
	if err != nil {
		return AddLocalCacheErr(err.Error())
	}
	return nil
}

// RemoveLocalCache linter
func RemoveLocalCache(key string) bool {
	return localCaching.Remove(key)
}

// GetLocalCache linter
func GetLocalCache(key string) (interface{}, error) {
	data, err := localCaching.Get(key)
	if err != nil {
		return nil, GetLocalCacheErr(err.Error())
	}
	return data, nil
}
