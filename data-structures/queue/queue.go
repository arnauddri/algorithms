package queue

type Queue struct {
	queue []interface{}
	len   int
}

func New() *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0

	return queue
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) isEmpty() bool {
	return q.len == 0
}

func (q *Queue) Shift() (el interface{}) {
	el, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return
}

func (q *Queue) Push(el interface{}) {
	q.queue = append(q.queue, el)
	q.len++
}

func (q *Queue) Peek() interface{} {
	return q.queue[0]
}
