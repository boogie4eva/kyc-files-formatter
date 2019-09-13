package model

import (
	"os"
	"sync"
)

type (
	// Worker must be implemented by types that want to use
	// the work pool.
	Worker interface {
		Task(info os.FileInfo)
	}

	// Pool provides a pool of goroutines that can execute any Worker
	// tasks that are submitted.
	Pool struct {
		work chan Worker
		wg   sync.WaitGroup
		file os.FileInfo
	}
)

// New creates a new work pool.
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task(p.file)
			}
			p.wg.Done()
		}()
	}
	return &p
}
