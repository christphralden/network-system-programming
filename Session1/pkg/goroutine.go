package pkg

import (
	"fmt"
	"sync"
	"time"
)

var rwm = sync.RWMutex{}
var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func RunGoroutine() {
	t0 := time.Now() // incriment counter
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait() // wait for counter to go back to 0
	fmt.Printf("Total Execution time:%v\n", time.Since(t0))

}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("result:", dbData[i])
	// rwm.RLock()
	m.Lock() //full lock
	results = append(results, dbData[i])
	// rwm.RUnlock()
	m.Unlock()
	wg.Done() //decrement counter
}
