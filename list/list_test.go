package list

import (
	"testing"
)

func TestNew(t *testing.T) {
	list := New(1, 2, 3)
	if list.Length() != 3 {
		t.Error("New: Expected length 3, got", list.Length())
	}

	emptyList := New[int]()
	if emptyList.Length() != 0 {
		t.Error("New: Expected length 0 for an empty list, got", emptyList.Length())
	}
}

func TestLength(t *testing.T) {
	list := New(1, 2, 3)
	if list.Length() != 3 {
		t.Error("Length: Expected length 3, got", list.Length())
	}

	emptyList := New[int]()
	if emptyList.Length() != 0 {
		t.Error("Length: Expected length 0 for an empty list, got", emptyList.Length())
	}
}

func TestPushBack(t *testing.T) {
	list := New(1, 2, 3)
	list.PushBack(4)
	if list.Length() != 4 {
		t.Error("PushBack: Expected length 4, got", list.Length())
	}

	val, _ := list.Get(3)
	if val != 4 {
		t.Error("PushBack: Expected value 4 at the end, got", val)
	}
}

func TestPushFront(t *testing.T) {
	list := New(1, 2, 3)
	list.PushFront(0)
	if list.Length() != 4 {
		t.Error("PushFront: Expected length 4, got", list.Length())
	}

	val, _ := list.Get(0)
	if val != 0 {
		t.Error("PushFront: Expected value 0 at the beginning, got", val)
	}
}

func TestInsert(t *testing.T) {
	list := New(1, 2, 3)
	list.Insert(4, 1)
	if list.Length() != 4 {
		t.Error("Insert: Expected length 4, got", list.Length())
	}

	val, _ := list.Get(1)
	if val != 4 {
		t.Error("Insert: Expected value 4 at index 1, got", val)
	}
}

func TestGet(t *testing.T) {
	list := New(1, 2, 3)
	val, err := list.Get(1)
	if err != nil || val != 2 {
		t.Error("Get: Expected value 2 at index 1, got", val, "with error", err)
	}

	_, err = list.Get(5)
	if err == nil {
		t.Error("Get: Expected out of range error, but got nil")
	}
}

func TestDelete(t *testing.T) {
	list := New(1, 2, 3)
	deleted, err := list.Delete(1)
	if err != nil || deleted != 2 {
		t.Error("Delete: Expected deleted value 2, got", deleted, "with error", err)
	}

	if list.Length() != 2 {
		t.Error("Delete: Expected length 2, got", list.Length())
	}
}

func TestPopFront(t *testing.T) {
	list := New(1, 2, 3)
	popped := list.PopFront()
	if popped != 1 {
		t.Error("PopFront: Expected popped value 1, got", popped)
	}

	if list.Length() != 2 {
		t.Error("PopFront: Expected length 2, got", list.Length())
	}
}

func TestPopBack(t *testing.T) {
	list := New(1, 2, 3)
	popped := list.PopBack()
	if popped != 3 {
		t.Error("PopBack: Expected popped value 3, got", popped)
	}

	if list.Length() != 2 {
		t.Error("PopBack: Expected length 2, got", list.Length())
	}
}

func TestClear(t *testing.T) {
	list := New(1, 2, 3)
	list.Clear()
	if list.Length() != 0 {
		t.Error("Clear: Expected length 0, got", list.Length())
	}
}
