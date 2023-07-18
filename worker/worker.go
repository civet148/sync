package worker

import (
	"errors"
	"sync"
)

type Workers struct {
	wg    sync.WaitGroup
	queue chan struct{}
}

// NewWorkers creates a Workers object which allows up to n jobs to proceed concurrently. n must be > 0.
func NewWorkers(n int) (*Workers, error) {
	if n <= 0 {
		return nil, errors.New("n must be > 0")
	}

	queue := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		queue <- struct{}{}
	}
	return &Workers{
		queue: queue,
	}, nil
}

// Take takes one from queue
func (w *Workers) Take() {
	w.wg.Add(1)
	<-w.queue
}

// Give returns one worker to queue
func (w *Workers) Give() {
	w.queue <- struct{}{}
	w.wg.Done()
}

// Wait waits for all ongoing concurrent jobs to complete
func (w *Workers) Wait() {
	w.wg.Wait()
}
