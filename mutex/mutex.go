package mutex

import (
	"sync"
	"fmt"
	"time"
)

func MutexNativeMeans() {
	receiver := make(chan bool)
	for i := 1; i < 10; i++ {
		for j:=1; j<10; j++ {
			go func () {
				fmt.Printf("Synced with sync.Mutex i is %d j is %d ", i, j)
				receiver <- true
			}()
			<- receiver
		}
	}
}

func MutexPackageUnSynced()  {
	pool := new(sync.Pool)
	pool.Get()
	fmt.Println("Time start: %s", time.Now())
	for i := 1; i < 10; i++ {
		for j:=1; j<10; j++ {
			go func () {
				fmt.Printf("Synced with sync.Mutex i is %d j is %d ", i, j)
			}()
		}
	}
}

func MutexPackageSync()  {
	pool := new(sync.Pool)
	pool.Get()
	mutex := new(sync.Mutex)
	fmt.Println("Time start: %s", time.Now())
	mutex.Lock()
	for i := 1; i < 10; i++ {
		for j:=1; j<10; j++ {
			go func () {
				mutex.Unlock()
				fmt.Printf("Synced with sync.Mutex i is %d j is %d ", i, j)
			}()
		}
	}
}