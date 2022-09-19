package mpool

import (
	"bytes"
	"testing"
)

func BenchmarkMPool(b *testing.B) {
	pools := NewMPool(100, func() interface{} {
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
