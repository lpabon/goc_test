package main

// #cgo CFLAGS: -pthread
// #include "pthread.h"
// #include "channel.h"
import "C"

import (
	"fmt"
	"sync"
	"time"
)

func channel_send_test(c_chan *C.channel_t, wg *sync.WaitGroup) {
	defer wg.Done()

	var pkt C.msg_t

	for i := 0; i < 50; i++ {
		C.ch_send(c_chan, &pkt)
		fmt.Println("PING >>")
		time.Sleep(time.Millisecond * time.Duration(i*2))
	}
}

func channel_recv_test(c_chan *C.channel_t, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 50; i++ {

		C.ch_recv(c_chan)
		fmt.Println("<< PONG")
		time.Sleep(time.Millisecond * time.Duration(i*5))
	}
}

func print_test(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Printf("%v: %v\n", name, i)
		time.Sleep(time.Millisecond * time.Duration(i))
	}
}

func main() {
	var wg sync.WaitGroup
	var c_chan C.channel_t

	C.ch_init(&c_chan)

	wg.Add(4)
	go channel_send_test(&c_chan, &wg)
	go channel_recv_test(&c_chan, &wg)
	go print_test("A", &wg)
	go print_test("B", &wg)

	wg.Wait()

}
