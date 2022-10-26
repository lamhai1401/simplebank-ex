package caching

import "time"

const (
	// DefaultExpireTime linter
	DefaultExpireTime = time.Duration(3600)
	// DefaultCapacity linter
	DefaultCapacity = 200
)

// Caching linter
type Caching interface {
	Set(key interface{}, value interface{}) error
	Get(key interface{}) (interface{}, error)
	Remove(key interface{}) bool
}
