package lru

import (
	"container/list"
)

type Cache struct {
	cap   int64                    //容量
	len   int64                    //已使用
	que   *list.List               //底层链表，此处当成双端队列来使用
	cache map[string]*list.Element //用于查找数据，真正的数据在list.Element中
}

type Value interface {
	Len() int64
}

type entry struct {
	key   string
	value Value //value有len属性，方便判断是否超出容量
}

//构造函数
func NewCache(len, cap int64) *Cache {
	return &Cache{
		cap:   cap,
		len:   len,
		que:   list.New(),
		cache: make(map[string]*list.Element),
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if v, ok := c.cache[key]; ok {
		c.que.MoveToFront(v)
		ent := v.Value.(*entry)
		return ent.value, true
	}
	return nil, false
}

//超出容量限制之后，删除lru链表中最后一个
func (c *Cache) Removelast() {
	v := c.que.Back()
	if v != nil {
		//队列中删除
		c.que.Remove(v)
		ent := v.Value.(*entry)
		//map中删除
		delete(c.cache, ent.key)
		//长度更新
		c.len -= int64(len(ent.key)) + ent.value.Len()
	}
}

func (c *Cache) Add(key string, value Value) {
	if v, ok := c.cache[key]; ok { //存在这个元素了
		c.que.MoveToFront(v)
		ent := v.Value.(*entry)
		//value修改之后可能带来的长度更新
		c.len += value.Len() - ent.value.Len()
		ent.value = value
	} else {
		ent := c.que.PushFront(&entry{key, value})
		c.cache[key] = ent
		c.len += int64(len(key)) + value.Len()
	}
	//超出容量限制
	for c.cap != 0 && c.cap < c.len {
		c.Removelast()
	}
}

func (c *Cache) Len() int {
	return c.que.Len()
}
