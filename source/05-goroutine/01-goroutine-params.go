package main

import (
	"fmt"
	"time"
)

func main() {
	animals := []string{"cat", "dog", "lion"}

	for _, v := range animals {
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(1 * time.Second)

	fmt.Println("\n")

	for _, v := range animals {
		go func(a string) {
			fmt.Println(a)
		}(v)
	}

	time.Sleep(1 * time.Second)
}

// output:

// lion
// lion
// lion

// lion
// cat
// dog
