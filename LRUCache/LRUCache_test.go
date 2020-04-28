package LRUCache

import (
	"reflect"
	"testing"
)

// LRUCache cache = new LRUCache( 2 /* capacity */ );

// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4

func TestConstructor(t *testing.T) {
	object := Constructor(2)

	objectType := reflect.TypeOf(object)
	LRUCacheType := reflect.TypeOf(LRUCache{})
	if objectType != LRUCacheType {
		t.Errorf("Constructor Error | Expected %v, got %v", LRUCacheType, objectType)
	}
}

func TestGet(t *testing.T) {
	object := Constructor(2)

	out := object.Get(0)
	if out != -1 {
		t.Errorf("Get Error | Expected: %d, got %v", -1, out)
	}

	object.Vals[0] = 1

	out = object.Get(0)
	if out != 1 {
		t.Errorf("Get Error | Expected: %d, got %v", 1, out)
	}
}

func TestPut(t *testing.T) {
	object := Constructor(2)

	object.Put(0, 1)

	out := object.Get(0)
	if out != 1 {
		t.Errorf("Get Error | Expected: %d, got %v", 1, out)
	}
}
