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
