package config

import (
	"github.com/LeyouHong/telo_demo/redis"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	MAP map[string]string
}

func newSafeMap() *SafeMap {
	sm := new(SafeMap)
	sm.MAP = make(map[string]string)
	return sm
}

func (sm *SafeMap) ReadMap(key string) string {
	sm.RLock()
	value := sm.MAP[key]
	sm.RUnlock()
	return value
}

func (sm *SafeMap) WriteMap(key string, value string) {
	sm.Lock()
	sm.MAP[key] = value
	sm.Unlock()
}

type Inputs struct {
	Names []string `json:"names"`
}

var Client = redis.NewRedis("redis:6379")

var Outputs = newSafeMap()