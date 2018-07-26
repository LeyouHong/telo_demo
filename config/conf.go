package config

import (
	"sync"
	sh "github.com/eaigner/shield"
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

type Shield struct {
	sync.RWMutex
	Shd sh.Shield
}

func newShield() *Shield {
	s := new(Shield)
	s.Shd = sh.New(
		sh.NewEnglishTokenizer(),
		sh.NewRedisStore("redis:6379", "",  nil, ""),
	)

	return s
}

func (s *Shield)Classify(val string) (string, error) {
	s.Lock()
	val, err := s.Shd.Classify(val)
	s.Unlock()
	return val, err
}

var Client = newShield()
var Outputs = newSafeMap()