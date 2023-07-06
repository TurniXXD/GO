package goroutines

import (
	"fmt"
	"sync"
	"time"
)

var (
	mtx        sync.Mutex
	rw_mtx     sync.RWMutex
	unProCount = 0
	proCount   = 0
)

func unprotectedIterator() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go unprotectedIncrement()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("Resulted count of unprotected iterator: ", unProCount)
}

func unprotectedIncrement() {
	unProCount++
}

func protectedIterator() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go protectedIncrement()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("Resulted count of mutex protected iterator: ", proCount)
}

// It protects resource so only one goroutine can access it at a time
func protectedIncrement() {
	mtx.Lock()
	defer mtx.Unlock()
	proCount++
}

func read() {
	rw_mtx.RLock()
	defer rw_mtx.RUnlock()

	fmt.Println("Read locking")
	time.Sleep(time.Second * 1)
	fmt.Println("Read unlocking")
}

func write(done chan bool) {
	rw_mtx.Lock()
	defer rw_mtx.Unlock()

	fmt.Println("Write locking")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Write unlocking")
	done <- true
}

// RWMutex stand for read-write mutual exclusion
func rw_mutex() {
	write_1_done := make(chan bool, 1)
	write_2_done := make(chan bool, 1)
	// RW mutex makes sure that multiple goroutines can read from protected resource at the same time but only one goroutine can write to protected resource at a time
	go read()
	go read()
	go write(write_1_done)
	go read()
	go write(write_2_done)
	go read()

	<-write_1_done
	<-write_2_done

	fmt.Println("RW done")
}

func rw_mutex_single_chan() {
	write_done := make(chan bool, 2)
	// RW mutex makes sure that multiple goroutines can read from protected resource at the same time but only one goroutine can write to protected resource at a time
	go read()
	go read()
	go write(write_done)
	go read()
	go write(write_done)
	go read()

	<-write_done
	<-write_done

	fmt.Println("RW done")
}

func protectedResource() {
	mtx.Lock()
	defer mtx.Unlock()

}

func syncMap() {
	// Maps are not safe for concurrent use
	// If multiple goroutines are accessing the same map, and at least one of them is writing to the map, the map must be locked with a mutex

	// If we run this we'll get fatal error "concurrent map writes"
	regularMap := make(map[int]interface{})

	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		regularMap[0] = i
	// 	}()
	// }

	// Instead we use sync.Map which provides safe concurrent access to the map
	syncMap := sync.Map{}

	for i := 0; i < 100; i++ {
		go func() {
			// Key, value
			syncMap.Store(0, i)
		}()
	}

	syncValue, ok := syncMap.Load(0)
	if ok {
		fmt.Println(syncValue)
	} else {
		fmt.Println("Could not load")
	}

	// Delete key 1
	syncMap.Delete(0)

	syncMap.Store(0, 6)

	// This
	syncValue, loaded := syncMap.LoadAndDelete(0)

	regularMap[0] = 7

	// Is same as this, but the syncMap is safer and is in the standard library
	mtx_map := sync.Mutex{}
	mtx_map.Lock()

	regularMapValue := regularMap[0]
	delete(regularMap, 0)

	mtx_map.Unlock()

	fmt.Println(syncValue, loaded, regularMapValue, regularMap[0])

}

// It stands for mutual exclusion
// It prevents race conditions, that happen when multiple goroutines are racing to access the same resource
func mutexes() {
	unprotectedIterator()
	protectedIterator()
	protectedResource()
	rw_mutex()
	// Faster
	rw_mutex_single_chan()
	syncMap()
}
