package pool

import (
	"log"
	"math"
	"sync"
)

type Task func() interface{}

type Pool struct {
	worker  int              // 并发数量
	tasks   *Queue           // 任务队列
	events  chan struct{}    // 任务通知管道
	Results chan interface{} // 结果管道
	wg      sync.WaitGroup   // 等待组
}

// 创建worker池
func NewPool(worker int) *Pool {
	return &Pool{
		worker:  worker,
		tasks:   NewQueue(-1),
		events:  make(chan struct{}, math.MaxInt32),
		Results: make(chan interface{}, worker*3),
	}
}

// 向任务池添加任务，并告诉通知管道
func (p *Pool) AddTask(task Task) {
	p.tasks.Append(task)
	// 通知任务池执行任务
	p.events <- struct{}{}
}

// 遍历通知管道，有通知则从队列中取出任务并执行
func (p *Pool) Start() {
	for i := 0; i < p.worker; i++ {
		p.wg.Add(1)
		go func(i int) {
			// 遍历通知管道
			for range p.events {
				// 从队列中读取任务
				e, err := p.tasks.Front()
				if err != nil {
					continue
				}

				log.Printf("worker %d run task", i)

				// 将队列中的空接口数据转换成Task并执行
				if task, ok := e.(Task); ok {
					// 将结果放入结果管道
					p.Results <- task()
				}
			}
			p.wg.Done()
		}(i)
	}
}

// 关闭任务池，等待任务执行完成
func (p *Pool) Close() {
	close(p.events)
	p.wg.Wait()
	close(p.Results)
}
