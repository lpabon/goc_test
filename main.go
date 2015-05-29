package main

import (
	"fmt"
	"sync"
	"time"
)

func print_test(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Printf("%v: %v\n", name, i)
		time.Sleep(time.Millisecond * time.Duration(i))
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go print_test("A", &wg)
	go print_test("B", &wg)

	wg.Wait()

}
