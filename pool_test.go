package mpool

import (
	"testing"

	"github.com/msaf1980/go-mpool/stack"
)

func TestNewPool(t *testing.T) {
	pool := NewPool(12, func() interface{} {
		return new(string)
	})

	if pool == nil {
		t.Errorf("NewPool is nil")
	}
}

func TestGetAndPut(t *testing.T) {
	var poolSize uint = 100
	pool := NewPool(poolSize, func() interface{} {
		return make([]int, 10)
	})

	poolNodes := make([]*stack.Node, 0, int(poolSize))

	for i := 0; i < int(poolSize); i++ {
		n := pool.Get()
		if n == nil {
			t.Errorf("Get is nil")
		}

		poolNodes = append(poolNodes, n)
	}

	expected := 0
	got := pool.Cap()

	if expected != got {
		t.Errorf("Cap expected: %v, got: %v", expected, got)
	}

	for _, n := range poolNodes {
		pool.Put(n)
	}

	expected = int(poolSize)
	got = pool.Cap()

	if expected != got {
		t.Errorf("Cap expected: %v, got: %v", expected, got)
	}
}

func TestCap(t *testing.T) {
	var poolSize uint = 10

	pool := NewPool(poolSize, func() interface{} {
		return new(int)
	})

	got := pool.Cap()

	if got != int(poolSize) {
		t.Errorf("Cap expected: %v, got: %v", poolSize, got)
	}

	_ = pool.Get()

	got = pool.Cap()

	if got != int(poolSize)-1 {
		t.Errorf("Cap expected: %v, got: %v", int(poolSize)-1, got)
	}
}

func TestDestPool(t *testing.T) {
	var poolSize uint = 20

	pool := NewPool(poolSize, func() interface{} {
		return make([]byte, 100)
	})

	expected := int(poolSize)
	got := pool.Cap()

	if expected != got {
		t.Errorf("Cap expected: %v, got: %v", expected, got)
	}

	pool.DestPool()

	expected = 0
	got = pool.Cap()

	if expected != got {
		t.Errorf("Cap expected: %v, got: %v", expected, got)
	}
}
