// Package memcachestore is a memcache based session store for the SCS session package.
package memcachestore

import (
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

type MemcacheStore struct {
	*memcache.Client
}

func New(server ...string) *MemcacheStore {
	return &MemcacheStore{
		memcache.New(server...),
	}
}

// search for session key
// returns value, found flag, error
func (p *MemcacheStore) Find(key string) ([]byte, bool, error) {
	item, err := p.Get(key)
	if err != nil {
		// session key not found
		if err == memcache.ErrCacheMiss {
			return nil, false, nil
		}

		// internal error
		return nil, false, err
	}

	return item.Value, true, nil
}

func (p *MemcacheStore) Save(key string, value []byte, expire time.Time) error {
	return p.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(expire.Unix()),
	})
}

func (p *MemcacheStore) Delete(token string) error {
	return p.Delete(token)
}
