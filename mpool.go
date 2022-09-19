package mpool

import (
	"github.com/msaf1980/go-mpool/stack"
)

// Pool is base Pool structor
type MPool struct {
	fn          func() interface{}
	stack       *stack.Stack
	initialSize int
}

// NewPool returns Pool instance
func NewMPool(uSize uint, fn func() interface{}) *MPool {
	iSize := int(uSize)

	s := stack.NewStack()

	for i := 0; i < iSize; i++ {
		s.Push(stack.NewNode(fn()))
	}

	return &MPool{
		stack: s,
		fn:    fn,
	}
}

// Get takes out of the pool
func (g *MPool) Get() *stack.Node {
	for {
		node, err := g.stack.Pop()
		if err != nil {
			g.upscale()
			continue
		}
		return node
	}
}

func (g *MPool) upscale() {
	for i := 0; i < g.initialSize; i++ {
		g.stack.Push(stack.NewNode(g.fn()))
	}
}

// Put puts node to pool
func (g *MPool) Put(node *stack.Node) {
	g.stack.Push(node)
}

// Cap returns current capacity of pool
func (g *MPool) Cap() int {
	return g.stack.Cap()
}

// DestPool destroys all pools
func (g *MPool) DestPool() {
	for !g.stack.IsEmpty() {
		g.stack.Pop()
	}
}
