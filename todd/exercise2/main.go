package main

import (
	"fmt"
	"sync/atomic"
	// "runtime"
	"sync"
)

func main() {
	// counter := 0
	var wg sync.WaitGroup
	var counter int64
	// var mu sync.Mutex

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			// mu.Lock()
			atomic.AddInt64(&counter, 1)
			// v := counter
			fmt.Println(atomic.LoadInt64(&counter))
			// runtime.Gosched()
			// v++
			// counter = v
			// mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	// fmt.Println(counter)
}

// const (
// 	b int = 43
// 	c float64 = 545.545
// )
// a := 42

// fmt.Printf("%d\n", a)
// fmt.Printf("%b\n", a)
// fmt.Printf("%#x\n", a)

// d := a << 1

// fmt.Printf("%d\n", d)
// fmt.Printf("%b\n", d)
// fmt.Printf("%#x\n", d)

// for i:= 0 ; i < 10001; i++ {
// 	fmt.Println(i)
// }

// x := []int{4,5,6,7,8,9}
// for _, v := range x {
// 	fmt.Println(v)
// }

// var wg sync.WaitGroup
// fmt.Println(runtime.NumGoroutine())
// wg.Add(2)

// go func() {

// 	fmt.Println("something")
// 	wg.Done()

// }()

// go func() {
// 	fmt.Println("something2")
// 	wg.Done()
// }()

// fmt.Println(runtime.NumGoroutine())
// wg.Wait()
// fmt.Println(runtime.NumGoroutine())
// fmt.Println("done here")
