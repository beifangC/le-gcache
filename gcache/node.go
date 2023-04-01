package gcache

import (
	"errors"
	"fmt"
	"log"
)

type Node struct {
	Name string
	C    cache
}

type Args struct {
	Key   string
	Value Byte
}

type Reply struct {
	Value Byte
}

func (n *Node) Add(args *Args, rep *Reply) error {
	n.add(args.Key, args.Value)
	return nil
}

func (n *Node) Get(args *Args, rep *Reply) error {
	value, err := n.get(args.Key)
	if err != nil {
		return fmt.Errorf("search fail")
	}
	rep.Value = value
	return nil
}

func (n *Node) get(key string) (Byte, error) {
	if key == "" {
		return Byte{}, fmt.Errorf("key is required")
	}

	if v, ok := n.C.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}
	return Byte{}, errors.New("not exist")
}

func (n *Node) add(key string, value Byte) error {
	if key == "" {
		return fmt.Errorf("key is required")
	}
	n.C.add(key, value)
	return nil
}
