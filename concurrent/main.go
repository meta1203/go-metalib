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

type WaitGroup struct {
	wg sync.WaitGroup
}

func NewWaitGroup() *WaitGroup {
	ret := WaitGroup{}
	ret.wg.Add(1)
	return &ret
}

func (WaitGroup *WaitGroup) Start() {
	WaitGroup.wg.Add(1)
}

func (WaitGroup *WaitGroup) Stop() {
	WaitGroup.wg.Add(1)
}

func (WaitGroup *WaitGroup) Wait() {
	WaitGroup.wg.Done()
	WaitGroup.wg.Wait()
}
