package Pool

import (
	"github.com/msaf1980/go-mpool/stack"
)

// Pool is nonthread-safe pool (for reuse allocated objects)
type Pool struct {
	fn          func() interface{}
	stack       *stack.Stack
	initialSize int
}

// NewPool returns Pool instance
func NewPool(uSize uint, fn func() interface{}) *Pool {
	iSize := int(uSize)

	s := stack.NewStack()

	for i := 0; i < iSize; i++ {
		s.Push(stack.NewNode(fn()))
	}

	return &Pool{
		stack: s,
		fn:    fn,
	}
}

// Get takes out of the pool
func (g *Pool) Get() *stack.Node {
	for {
		node, err := g.stack.Pop()
		if err != nil {
			g.upscale()
			continue
		}
		return node
	}
}

func (g *Pool) upscale() {
	for i := 0; i < g.initialSize; i++ {
		g.stack.Push(stack.NewNode(g.fn()))
	}
}

// Put puts node to pool
func (g *Pool) Put(node *stack.Node) {
	g.stack.Push(node)
}

// Cap returns current capacity of pool
func (g *Pool) Cap() int {
	return g.stack.Cap()
}

// DestPool destroys all pools
func (g *Pool) DestPool() {
	for !g.stack.IsEmpty() {
		g.stack.Pop()
	}
}
