package worker

import (
	"context"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type Task interface {
	Execute() error
}

type Pool struct {
	TaskChan    chan Task
	Concurrency int

	Workers map[string]worker
	mut     sync.RWMutex

	addWorkerChan    chan struct{}
	cancelWorkerChan chan struct{}

	ctx    context.Context
	cancel context.CancelFunc
}

func NewPool(ctx context.Context, concurrent int) *Pool {
	ctx, cancel := context.WithCancel(ctx)

	p := &Pool{
		TaskChan:         make(chan Task),
		addWorkerChan:    make(chan struct{}),
		cancelWorkerChan: make(chan struct{}),
		Concurrency:      concurrent,
		Workers:          make(map[string]worker),
		ctx:              ctx,
		cancel:           cancel,
	}

	for i := 0; i < p.Concurrency; i++ {
		p.addWorker()
	}

	go func() {
		for {
			select {
			case <-p.addWorkerChan:
				p.addWorker()

			case <-p.cancelWorkerChan:
				p.delOnceWorker()

			case <-time.NewTicker(time.Second * 1).C:
				if len(p.Workers) < p.Concurrency {
					go func() {
						p.addWorkerChan <- struct{}{}
					}()
				}

				if len(p.Workers) > p.Concurrency {
					go func() {
						p.cancelWorkerChan <- struct{}{}
					}()
				}
			case <-p.ctx.Done():
				return
			}
		}
	}()

	return p
}

func (p *Pool) addWorker() {
	w := newWorker(p.ctx, p.TaskChan)
	go w.Run()
	log.Infof("new worker [%s]", w.ID)

	p.mut.Lock()
	p.Workers[w.ID] = *w
	p.mut.Unlock()

	log.Infof("add worker [%s] to pool", w.ID)
}

func (p *Pool) delOnceWorker() {
	var worker worker

	p.mut.RLock()
	for _, w := range p.Workers {
		worker = w
		break
	}
	p.mut.RUnlock()

	log.Infof("cancel worker [%s]", worker.ID)
	worker.Cancel()

	p.mut.Lock()
	delete(p.Workers, worker.ID)
	p.mut.Unlock()

	log.Infof("delete worker [%s] from pool", worker.ID)
}

func (p *Pool) Submit(task Task) {
	p.TaskChan <- task
}

func (p *Pool) SetPoolSize(c int) {
	p.Concurrency = c
}

type worker struct {
	ID       string
	taskChan chan Task

	ctx    context.Context
	Cancel context.CancelFunc
}

func newWorker(ctx context.Context, taskChan chan Task) *worker {
	c, cancel := context.WithCancel(ctx)

	return &worker{
		ID:       uuid.NewV4().String(),
		taskChan: taskChan,

		ctx:    c,
		Cancel: cancel,
	}
}

func (w *worker) Run() {
	for {
		select {
		case task := <-w.taskChan:
			log.Infof("worker [%s] receive task", w.ID)
			w.run(task)
		case <-w.ctx.Done():
			log.Infof("worker [%s] exit", w.ID)
			return
		}
	}
}

func (w *worker) run(t Task) {
	var err error

	defer func() {
		if err != nil {
			log.Errorf("task error, msg: %v", err)
		}
	}()

	err = t.Execute()
	log.Infof("task execute completed")
}

func (w *worker) Submit(t Task) {
	w.taskChan <- t
}
