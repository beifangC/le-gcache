package main

import (
	"fmt"
	"gcache/gcache"
	"gcache/myrpc"
	"log"
	_ "log"
	_ "net"
	_ "net/http"
	_ "net/rpc"
)

// 服务端
/*
func main() {
	node := &myrpc.Node{
		G: *gcache.NewGroup("node", 2<<10),
	}
	rpc.Register(node)
	rpc.HandleHTTP()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error", err)
	}
	http.Serve(lis, nil)
}
*/
//客户端
func main() {
	a := myrpc.Rpcclient{}
	v := gcache.Byte{
		[]byte("test"),
	}
	err := a.Add("tom", v)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := a.Get("tom")
	fmt.Println(b)
}
