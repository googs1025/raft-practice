package cache

import (
	"sync"
)

var (
	once sync.Once
	m    *sync.Map
)

type MapCache struct {
	*sync.Map
}

func NewMapCache() Cache {
	once.Do(func() {
		m = &sync.Map{}
	})
	return &MapCache{m}
}

func GetCache(mc *MapCache) *sync.Map {
	return mc.Map
}

func (m *MapCache) SetItem(key string, value string) error {
	m.Store(key, value)
	return nil
}

func (m *MapCache) GetItem(key string) (string, error) {
	if v, ok := m.Load(key); ok {
		res, o := v.(string)
		if !o {
			return "", nil
		}
		return res, nil
	}
	return "", nil
}

func (m *MapCache) DelItem(key string) error {
	m.Delete(key)
	return nil
}
