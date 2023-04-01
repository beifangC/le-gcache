package gcache

//对lru进行封装

type Byte struct {
	B []byte
}

func (v Byte) Len() int64 {
	return int64(len(v.B))
}

func (v Byte) ByteSlice() []byte {
	return cloneBytes(v.B)
}

func (v Byte) String() string {
	return string(v.B)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
