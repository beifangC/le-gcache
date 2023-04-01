package myrpc

import (
	"errors"
	"gcache/gcache"
)

type Node struct {
	G gcache.Group
}

type Args struct {
	Key   string
	Value gcache.Byte
}

type Reply struct {
	Value gcache.Byte
}

func (n *Node) Add(args *Args, rep *Reply) error {
	n.G.Add(args.Key, args.Value)
	return nil
}

func (n *Node) Get(args *Args, rep *Reply) error {
	res, err := n.G.Get(args.Key)
	if err != nil {
		return errors.New("not exist")
	}
	rep.Value = res
	return nil
}
