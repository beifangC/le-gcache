package gcache

import (
	"fmt"
	"log"
	"sync"
)

//实现单机的缓存
type Group struct {
	name string
	ca   cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cap int64) *Group {
	mu.Lock()
	defer mu.Unlock()
	group := &Group{
		name: name,
		ca:   cache{cap: cap},
	}
	//加入group集群
	groups[name] = group
	return group
}

func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

// Get value by key
func (g *Group) Get(key string) (Byte, error) {
	if key == "" {
		return Byte{}, fmt.Errorf("key is required")
	}

	if v, ok := g.ca.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}

	return g.load(key)
}

func (g *Group) load(key string) (value Byte, err error) {
	return
}
