package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex

func testMap() {
	var a map[int]int
	a = make(map[int]int, 5)

	a[1] = 2
	a[2] = 3

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[1] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
	time.Sleep(time.Second)
}

var rwlock sync.RWMutex

func testRWLock() {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32

	a[1] = 2
	a[2] = 3

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			// rwlock.Lock()
			lock.Lock()
			b[1] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			// rwlock.Unlock()
			lock.Unlock()
		}(a)
	}

	for i := 0; i < 100000; i++ {
		go func(b map[int]int) {
			for {
				lock.Lock()
				// rwlock.RLock()
				//fmt.Println(a)
				time.Sleep(time.Millisecond)
				// rwlock.RUnlock()
				lock.Unlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}

	time.Sleep(time.Second * 5)
	fmt.Println(atomic.LoadInt32(&count))
}

func main() {
	// testMap()
	testRWLock()
}
