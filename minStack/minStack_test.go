package minStack

import (
	"reflect"
	"testing"
)

// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> Returns -3.
// minStack.pop();
// minStack.top();      --> Returns 0.
// minStack.getMin();   --> Returns -2.

func TestConstructor(t *testing.T) {
	object := Constructor()
	objectType := reflect.TypeOf(object)
	minStackType := reflect.TypeOf(MinStack{})
	if objectType != minStackType {
		t.Errorf("Constructor Error | Expected %v, got %v", minStackType, objectType)
	}
}

func TestPush(t *testing.T) {

	// Instantiate stack
	object := Constructor()

	// Check empty stack
	length := len(object.stack)
	if length != 0 {
		t.Errorf("Push Error | Object length. Expected: %v, got %v", 0, length)
		return
	}

	// Push one value; check stack length == 1
	object.Push(3)
	length = len(object.stack)
	if length != 1 {
		t.Errorf("Push Error | Object length. Expected: %v, got %v", 1, length)
		return
	}

	// Check value matches pushed value
	got := object.stack[0]
	if got != 3 {
		t.Errorf("Push Error | Object value. Expected %v, got %v", 3, got)
		return
	}

}

func TestPop(t *testing.T) {
	// Instantiate stack
	object := Constructor()

	// Check empty stack
	length := len(object.stack)
	if length != 0 {
		t.Errorf("Pop Error| Object length. Expected: %v, got %v", 0, length)
		return
	}

	// Push one value; check stack length == 1
	object.Push(3)
	length = len(object.stack)
	if length != 1 {
		t.Errorf("Pop Error | Object length. Expected: %v, got %v", 1, length)
		return
	}

	// Pop first value; check stack length == 0
	object.Pop()
	length = len(object.stack)
	if length != 0 {
		t.Errorf("Pop Error | Object length. Expected: %v, got %v", 0, length)
		return
	}

}

func TestTop(t *testing.T) {
	// Instantiate stack
	object := Constructor()

	// Check empty stack
	length := len(object.stack)
	if length != 0 {
		t.Errorf("Top Error | Object length. Expected: %v, got %v", 0, length)
		return
	}

	// Push two values
	object.Push(3)
	object.Push(10)

	max := object.Top()
	if max != 10 {
		t.Errorf("Top Error | Max. Expected: %v, got: %v", 10, max)
	}
}

func TestGetMin(t *testing.T) {
	// Instantiate stack
	object := Constructor()

	// Check empty stack
	length := len(object.stack)
	if length != 0 {
		t.Errorf("Top Error | Object length. Expected: %v, got %v", 0, length)
		return
	}

	// Push two values
	object.Push(3)
	object.Push(10)

	max := object.GetMin()
	if max != 3 {
		t.Errorf("Top Error | Min. Expected: %v, got: %v", 3, max)
	}

}
