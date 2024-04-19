package generics

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type node[T any] struct {
	next *node[T]
	val  T
}

type List[T any] struct {
	head, tail *node[T]
}

func (list *List[T]) Push(val T) {
	if list.tail == nil {
		list.tail = &node[T]{val: val}
		list.head = list.tail
	} else {
		list.tail.next = &node[T]{val: val}
		list.tail = list.tail.next
	}
}

func (list List[T]) GetAll() []T {
	var elements []T

	for node := list.head; node != nil; node = node.next {
		elements = append(elements, node.val)
	}
	return elements
}

func Max[A int](x, y A) A {
	if x > y {
		return x
	}
	return y
}
