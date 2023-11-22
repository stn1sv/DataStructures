package list

import "fmt"

// node - узел связного списка
type node[T comparable] struct {
	data T
	next *node[T]
}

// List - фиктивный (головной) узел (элемент) связного списка
type List[T comparable] struct {
	head   *node[T] //указатель на первый элемент списка
	tail   *node[T] //указатель на последний элемент списка
	length int
}

// New создает связный список и добавляет в него переданные аргументы
func New[T comparable](elems ...T) *List[T] {
	list := &List[T]{length: len(elems)}
	if list.length == 0 {
		return list
	}

	list.head = &node[T]{data: elems[0]}
	current := list.head
	for i := 1; i < len(elems); i++ {
		current.next = &node[T]{data: elems[i]}
		current = current.next
	}
	list.tail = current

	return list
}

// Length возвращает длину связного списка
// Асимптотическая сложность - О(1)
func (l *List[T]) Length() int {
	return l.length
}

// Get возвращает элемент по заданному индексу
// Асимптотическая сложность - О(n)
func (l *List[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.length {
		var v T
		return v, fmt.Errorf("index %d is out of range: list length is %d", index, l.length)
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.data, nil
}

// Print выводит связный список
// Асимптотическая сложность - О(n)
func (l *List[T]) Print() {
	if l.length == 0 {
		fmt.Println("Список пуст")
		return
	}

	fmt.Print("[")
	for current := l.head; ; current = current.next {
		if current.next == nil {
			fmt.Print(current.data, "]\n")
			return
		} else {
			fmt.Print(current.data, ", ")
		}
	}
}

// PushBack добавляет элемент в конец списка
// Асимптотическая сложность - О(1)
func (l *List[T]) PushBack(data T) *List[T] {
	n := &node[T]{data: data}

	if l.length == 0 {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
	l.length++

	return l
}

// PushFront добавляет элемент в начало списка
// Асимптотическая сложность - О(1)
func (l *List[T]) PushFront(data T) *List[T] {
	n := &node[T]{data: data, next: l.head}

	l.head = n
	if l.length == 0 {
		l.tail = n
	}
	l.length++

	return l
}

// Insert добавляет элемент в середину списка
// Асимптотическая сложность - О(n)
func (l *List[T]) Insert(data T, index int) (*List[T], error) {
	if index < 0 || index > l.length {
		return nil, fmt.Errorf("index %d is out of range: list length is %d", index, l.length)
	}

	if index == 0 {
		return l.PushFront(data), nil
	}

	if index == l.length {
		return l.PushBack(data), nil
	}

	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	n := &node[T]{data: data, next: current.next}
	current.next = n
	l.length++

	return l, nil
}

// Delete удаляеет элемент по заданному индексу
// Асимптотическая сложность - О(n)
func (l *List[T]) Delete(index int) (T, error) {
	if index < 0 || index > l.length {
		var v T
		return v, fmt.Errorf("index %d is out of range: list length is %d", index, l.length)
	}

	if index == 0 {
		return l.PopFront(), nil
	}

	prev := l.head
	current := prev.next
	for i := 1; i < index; i++ {
		prev = current
		current = current.next
	}
	prev.next = current.next
	if index == l.length-1 {
		l.tail = prev
	}
	l.length--

	return current.data, nil
}

// PopFront удаляет первый (нулевой по индексу) элемент списка
// Асимптотическая сложность - О(1)
func (l *List[T]) PopFront() T {
	result := l.head.data
	l.head = l.head.next
	l.length--

	return result
}

// PopBack удаляет последний элемент
// Асимптотическая сложность - О(n)
func (l *List[T]) PopBack() T {
	if l.head == nil {
		var v T
		return v
	}

	if l.head.next == nil {
		v := l.head.data
		l.head = nil
		l.tail = nil
		l.length--
		return v
	}

	prev := l.head
	current := prev.next
	for current.next != nil {
		prev = current
		current = current.next
	}
	prev.next = nil
	l.tail = prev
	l.length--

	return current.data
}

// Clear отчищает список
func (l *List[T]) Clear() *List[T] {
	l.head = nil
	l.tail = nil
	l.length = 0

	return l
}
