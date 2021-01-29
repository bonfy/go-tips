package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("start worker %d\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("quit worker %d\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}

// output:

// start worker 4
// start worker 1
// start worker 0
// start worker 3
// start worker 2
// quit worker 4
// quit worker 0
// quit worker 1
// quit worker 2
// quit worker 3
