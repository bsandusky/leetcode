package firstUnique

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	object := Constructor([]int{0, 1, 2, 3, 4})

	objectType := reflect.TypeOf(object)
	FirstUniqueType := reflect.TypeOf(FirstUnique{})
	if objectType != FirstUniqueType {
		t.Errorf("Constructor Error | Expected %v, got %v", FirstUniqueType, objectType)
	}
}

func TestShowFirstUnique(t *testing.T) {
	object := Constructor([]int{2, 3, 5})
	out := object.ShowFirstUnique()

	if out != 2 {
		t.Errorf("Show First Unique Error 1 | Expected value of %d, got %d", 2, out)
	}

	object.Add(5)
	out = object.ShowFirstUnique()

	if out != 2 {
		t.Errorf("Show First Unique Error 2 | Expected value of %d, got %d", 2, out)
	}

	object.Add(2)
	out = object.ShowFirstUnique()

	if out != 3 {
		t.Errorf("Show First Unique Error 3 | Expected value of %d, got %d", 3, out)
	}

	object.Add(3)
	out = object.ShowFirstUnique()

	if out != -1 {
		t.Errorf("Show First Unique Error 4 | Expected value of %d, got %d", -1, out)
	}
}

func TestAnotherShowFirstUnique(t *testing.T) {
	object := Constructor([]int{7, 7, 7, 7, 7, 7})
	out := object.ShowFirstUnique()

	if out != -1 {
		t.Errorf("Show First Unique Error 1 | Expected value of %d, got %d", -1, out)
	}

	object.Add(7)
	object.Add(3)
	object.Add(3)
	object.Add(7)
	object.Add(17)

	out = object.ShowFirstUnique()

	if out != 17 {
		t.Errorf("Show First Unique Error 1 | Expected value of %d, got %d", 17, out)
	}

}

func TestAdd(t *testing.T) {
	object := Constructor([]int{0, 1, 2, 3, 4})

	if len(object.Counts) != 5 {
		t.Errorf("Add Error | Expected length of %d, got %d", 5, len(object.Counts))
	}

	object.Add(5)

	if len(object.Counts) != 6 {
		t.Errorf("Add Error | Expected length of %d, got %d", 6, len(object.Counts))
	}

	if object.Counts[len(object.Counts)-1] != 5 {
		t.Errorf("Constructor Error | Expected value of %d, got %d", 5, len(object.Counts))
	}
}
