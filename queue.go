package ds

type Queue[T any] struct {
	buffer []T
	head   int
	tail   int
	size   int
}

func NewQueue[T any](options ...int) Queue[T] {
	size := 8
	if len(options) > 0 {
		size = options[0]
	}

	// This implementation can hold up to n-1 elements
	size++

	return Queue[T]{
		buffer: make([]T, size),
		head:   0,
		tail:   0,
		size:   size,
	}
}

func (q *Queue[T]) Enqueue(elem T) {
	if q.isFull() {
		q.grow()
	}

	q.buffer[q.tail] = elem
	q.tail = (q.tail + 1) % q.size
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var out T

	if q.isEmpty() {
		return out, false
	}

	out = q.buffer[q.head]
	q.head = (q.head + 1) % q.size

	return out, true
}

func (q Queue[T]) isFull() bool {
	return (q.tail+1)%q.size == q.head
}

func (q Queue[T]) isEmpty() bool {
	return q.head == q.tail
}

func (q *Queue[T]) grow() {
	newSize := (q.size-1)*2 + 1
	newBuffer := make([]T, newSize)
	copy(newBuffer, q.buffer)

	q.buffer = newBuffer
	q.size = newSize
}
