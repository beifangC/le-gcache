package myrpc

import (
	"fmt"
	"gcache/consishash"
	"gcache/gcache"
	"net/rpc"
	"os"

	"gopkg.in/ini.v1"
)

var Clients map[string]*rpc.Client
var Chash *consishash.Map

type Rpcclient struct{}

func init() {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	names := config.SectionStrings() //[DEFAULT mysql node3 node1 node0]
	//fmt.Println(names)
	//创建哈希环
	Chash = consishash.New(10, nil)

	//保存节点对应的tcp链接
	Clients = make(map[string]*rpc.Client, len(names)-2)
	for i := 2; i < len(names); i++ {
		n := names[i]
		ip := config.Section(names[i]).Key("ip").String()
		client, _ := rpc.DialHTTP("tcp", ip)
		Clients[n] = client
		//添加节点到哈希环
		Chash.Add(n)
	}

}

func (c *Rpcclient) Add(key string, value gcache.Byte) error {
	node := Chash.Get(key) //获取节点名称
	link := Clients[node]  //获取对应链接
	args := &Args{
		Key:   key,
		Value: value,
	}
	rep := Reply{}
	err := link.Call("Node.Add", args, &rep)
	if err != nil {
		return err
	}
	return nil
}

func (c *Rpcclient) Get(key string) (gcache.Byte, error) {
	node := Chash.Get(key) //获取节点名称
	link := Clients[node]  //获取对应链接
	args := &Args{
		Key: key,
	}
	rep := Reply{}
	err := link.Call("Node.Get", args, &rep)
	if err != nil {
		return gcache.Byte{}, err
	}
	return rep.Value, nil
}
