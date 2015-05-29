package main

// #cgo CFLAGS: -pthread
// #include "pthread.h"
// #include "channel.h"
import "C"

import (
	"fmt"
	"github.com/lpabon/godbc"
	"sync"
	"time"
)

func channel_send_test(c_chan *C.channel_t, wg *sync.WaitGroup) {
	defer wg.Done()

	var pkt C.msg_t

	for i := 0; i < 50; i++ {
		pkt.val = C.int(i)
		fmt.Printf("PING >> %v\n", pkt.val)
		C.ch_send(c_chan, &pkt)
		time.Sleep(time.Millisecond * time.Duration(i*2))
	}
}

func channel_recv_test(c_chan *C.channel_t, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 50; i++ {

		pkt := C.ch_recv(c_chan)
		godbc.Check(C.int(i) == pkt.val, i, int(pkt.val))
		fmt.Printf("%v << PONG\n", pkt.val)
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
