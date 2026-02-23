package dsa

type Queue[T any] struct {
	head int
	buf  []T
}

func (q *Queue[T]) Push(v T) {
	q.buf = append(q.buf, v)
}

func (q *Queue[T]) Pop() bool {
	if q.head == len(q.buf) {
		return false
	}
	q.head += 1

	// compaction
	if q.head > 64 && q.head*2 > len(q.buf) {
		n := copy(q.buf, q.buf[q.head:])
		q.buf = q.buf[:n]
		q.head = 0
	}
	return true
}

func (q *Queue[T]) Front() T {
	return q.buf[q.head]
}

func (q *Queue[T]) Empty() bool {
	return q.head == len(q.buf)
}

func (q *Queue[T]) Size() int {
	return len(q.buf) - q.head
}
