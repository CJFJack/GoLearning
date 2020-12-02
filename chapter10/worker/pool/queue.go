package pool

import (
	"fmt"
	"sync"
)

const defaultCap = 1024

type Queue struct {
	elements []interface{}
	locker   sync.Mutex
	limit    int
}

// 创建队列
func NewQueue(limit int) *Queue {
	return &Queue{
		elements: make([]interface{}, 0, defaultCap),
		locker:   sync.Mutex{},
		limit:    limit,
	}
}

// 队列末尾加入元素
func (q *Queue) Append(e interface{}) error {
	if q.limit != -1 && len(q.elements) >= q.limit {
		return fmt.Errorf("queue is full, limit is %d", q.limit)
	}
	q.locker.Lock()
	defer q.locker.Unlock()
	q.elements = append(q.elements, e)
	return nil
}

// 取出第一个元素
func (q *Queue) Front() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, fmt.Errorf("queue is nil")
	}
	q.locker.Lock()
	defer q.locker.Unlock()
	e := q.elements[0]
	q.elements = q.elements[1:]
	return e, nil
}

func (q *Queue) Len() int {
	return len(q.elements)
}
