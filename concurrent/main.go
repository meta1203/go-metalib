package concurrent

import (
	"log"
	"os"
	"sync"
)

// a concurrency-safe stdout printer
var Log *log.Logger = log.New(os.Stdout, "", 0)

// create a mutex lock
func CreateLock() *sync.Mutex {
	return &sync.Mutex{}
}

type waitGroup struct {
	wg sync.WaitGroup
}

func NewWaitGroup() *waitGroup {
	ret := WaitGroup{}
	ret.wg.Add(1)
	return &ret
}

func (waitGroup *waitGroup) Start() {
	waitGroup.wg.Add(1)
}

func (waitGroup *waitGroup) Stop() {
	waitGroup.wg.Add(1)
}

func (waitGroup *waitGroup) Wait() {
	waitGroup.wg.Done()
	waitGroup.wg.Wait()
}
