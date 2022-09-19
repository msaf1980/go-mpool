package mpool

import (
	"bytes"
	"sync"
	"testing"
)

func BenchmarkPool(b *testing.B) {
	pools := NewPool(100, func() interface{} {
		return &bytes.Buffer{}
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool1 := pools.Get()
		_ = pool1.Value.(*bytes.Buffer)

		pool2 := pools.Get()
		_ = pool2.Value.(*bytes.Buffer)

		pools.Put(pool1)
		pools.Put(pool2)
	}
}

func BenchmarkSyncPool(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf1 := pool.Get().(*bytes.Buffer)

		buf2 := pool.Get().(*bytes.Buffer)

		pool.Put(buf1)
		pool.Put(buf2)
	}
}
