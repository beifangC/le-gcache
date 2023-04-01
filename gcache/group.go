package gcache

import (
	"errors"
	"fmt"
	"log"
	"sync"
)
//单机版本

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

//
func (g *Group) Get(key string) (Byte, error) {
	if key == "" {
		return Byte{}, fmt.Errorf("key is required")
	}

	if v, ok := g.ca.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}

	return Byte{}, errors.New("not exist")
}

func (g *Group) Add(key string, value Byte) error {
	if key == "" {
		return fmt.Errorf("key is required")
	}
	g.ca.add(key, value)
	return nil
}
