package stack

import (
	"runtime"
	"sync/atomic"
)

// The Treiber stack algorithm is a scalable lock-free stack utilizing the fine-grained concurrency primitive compare-and-swap.
// [1] It is believed that R. Kent Treiber was the first to publish it in his 1986 article "Systems Programming: Coping with Parallelism"

// Stack is treiber stack
type MStack struct {
	head *Node
	lf   *int32 // lock flag
}

// NewStack returns stack instance
func NewMStack() *MStack {
	return &MStack{
		head: nil,
		lf:   new(int32),
	}
}

func wait(lf *int32) {
	for {
		if atomic.CompareAndSwapInt32(lf, 0, 1) {
			break
		}
		// time.Sleep(1 * time.Nanosecond)
		runtime.Gosched()
	}
}

// Signal signals termination of use of shared counter
func signal(lf *int32) {
	atomic.StoreInt32(lf, 0)
}

// Push appends value into the stack
func (s *MStack) Push(newHead *Node) {
	wait(s.lf)

	newHead.next = s.head
	s.head = newHead

	signal(s.lf)
}

// Pop returns node of the stack
func (s *MStack) Pop() (*Node, error) {
	wait(s.lf)

	if s.head == nil {
		return nil, ErrStackEmpty
	}

	tmpHead := s.head
	s.head = tmpHead.next

	signal(s.lf)

	return tmpHead, nil
}

// IsEmpty returns true if the stack is empty, one the other hand, it returns false if it is not empty
func (s *MStack) IsEmpty() bool {
	wait(s.lf)
	empty := s.head == nil
	signal(s.lf)
	return empty
}

// Cap returns current capacity of stack
func (s *MStack) Cap() (cnt int) {
	wait(s.lf)

	tmpHead := s.head
	for tmpHead != nil {
		cnt++
		tmpHead = tmpHead.next
	}

	signal(s.lf)

	return cnt
}
