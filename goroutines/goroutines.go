package goroutines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var result = []string{}

func Goroutines() {
	fmt.Println("Hello ! Welcome to world of World of GoRoutines where we talk about concurrency and parallel executions.")
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nThe total executation time: %v ", time.Since(t0))
	fmt.Printf("\nThe result is: %v\n", result)
}

//func count() {
//	var res int
//	for i := 0; i < 100000000; i++ {
//		res += 1
//	}
//	//wg.Done()
//}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is : ", dbData[i])
	mutex.Lock()
	result = append(result, dbData[i])
	mutex.Unlock()
	wg.Done()
}
