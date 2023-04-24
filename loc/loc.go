package loc

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

//保证key对应的fn只被执行一次
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok { //请求已经在进行
		g.mu.Unlock()
		c.wg.Wait() //等待返回
		return c.val, c.err
	}
	c := new(call) //请求没有进行
	c.wg.Add(1)
	g.m[key] = c //标记
	g.mu.Unlock()

	c.val, c.err = fn() //新的请求
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key) //运行结束，删除标记
	g.mu.Unlock()

	return c.val, c.err
}
