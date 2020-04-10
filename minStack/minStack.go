package minStack

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// MinStack type
type MinStack struct {
	stack []int
	min   []int
}

// Constructor : initializes your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

// Push appends item to stack array
func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)
	if len(m.min) == 0 || m.GetMin() >= x {
		m.min = append(m.min, x)
	}
}

// Pop removes first element of stack array
func (m *MinStack) Pop() {
	if len(m.stack) == 0 {
		return
	}

	if m.Top() == m.GetMin() {
		m.min = m.min[:len(m.min)-1]
	}

	m.stack = m.stack[:len(m.stack)-1]
}

// Top returns top element of stack array
func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

// GetMin returns lowest value within stack array
func (m *MinStack) GetMin() int {
	return m.min[len(m.min)-1]
}
