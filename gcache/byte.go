package gcache

//对lru进行封装，是Value的一种实现
type Byte struct {
	b []byte
}

// Len returns the view's length
func (v Byte) Len() int64 {
	return int64(len(v.b))
}

// ByteSlice returns a copy of the data as a byte slice.
func (v Byte) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary.
func (v Byte) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
