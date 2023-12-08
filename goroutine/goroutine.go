package goroutine

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
	"sync"
)

func SafeGo(ctx context.Context, f func()) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				log.Printf("safe goroutine panic, e:%v %s", e, string(debug.Stack()))
			}
		}()
		select {
		case <-ctx.Done():
			return
		default:
			f()
		}
	}()
}

func printOddAndEven(start, end int) {
	odd := make(chan int)
	even := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for val := range odd {
			fmt.Println(val)
			if val == end {
				close(even)
				break
			}
			even <- val + 1
		}
		wg.Done()
	}()
	go func() {
		for val := range even {
			fmt.Println(val)
			if val == end {
				close(odd)
				break
			}
			odd <- val + 1
		}
		wg.Done()
	}()
	odd <- start
	wg.Wait()
}
