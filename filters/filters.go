package filters

import (
	"fmt"
	"runtime"
)

func InitFilters() {
	runtime.GOMAXPROCS(4)
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Println(prime)
		chOut := make(chan int)
		go filter (ch, chOut, prime)
		ch = chOut
	}
}

func generate (ch chan int) {
	for i := 2;; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int)  {
	for {
		i := <-in
		//any filter statements can be here
		if i%prime != 0 {
			out <- i
		}
	}
}