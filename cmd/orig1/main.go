package main

import (
	"fmt"
	"sync"
	"time"
)

type Hoge interface {
	Right() string
}

type UserNotFound struct {
	Username string
}

func (u *UserNotFound) Error() string {
	return fmt.Sprintf("not found %v", u.Username)
}

func (u *UserNotFound) Right() string {
	return u.Username
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	hoge := &UserNotFound{Username: "mike"}
	fmt.Println(hoge.Right())

	return hoge
}

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			fmt.Println("process ", i*1000)
			defer wg.Done()
		}()

	}
}

func producer2(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multi4(second chan int, third chan int) {
	defer close(third)
	for i := range second {
		third <- i * 2
	}
}

func goroutine1(ch chan<- string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutine2(ch chan<- string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]

}

// make new 違い
// ポインタを返すか返さないか
func main() {
	var p *int = new(int) // アドレス確保
	fmt.Println(*p)

	var p2 *int
	fmt.Println(p2)

	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("aaa")

	//ch := make(chan int, 2) // buffer数
	//ch <- 100
	//fmt.Println(len(ch))
	//ch <- 200
	//fmt.Println(len(ch))
	//close(ch)
	//
	//for c := range ch {
	//    fmt.Println(c)
	//}

	//// pro, consumer
	//var wg sync.WaitGroup
	//ch := make(chan int)
	//
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go producer(ch, i)
	//}
	//
	//go consumer(ch, &wg)
	//wg.Wait()
	//close(ch)

	//// fan-out fan-in
	//first := make(chan int)
	//second := make(chan int)
	//third := make(chan int)
	//
	//go producer2(first)
	//go multi2(first, second)
	//go multi4(second, third)
	//for result := range third {
	//	fmt.Println(result)
	//}

	//// chan select
	//ch1 := make(chan string)
	//ch2 := make(chan string)
	//go goroutine1(ch1)
	//go goroutine2(ch2)
	//
	//for {
	//	select {
	//	case msg1 := <-ch1:
	//		fmt.Println(msg1)
	//	case msg2 := <-ch2:
	//		fmt.Println(msg2)
	//	}
	//}

	//// default selection
	//tick := time.Tick(100 * time.Millisecond)
	//boom := time.After(500 * time.Millisecond)
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick.")
	//	case <-boom:
	//		fmt.Println("BOOM!")
	//		return
	//	default:
	//		fmt.Println("    .")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}

	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c.Value("key"))
}
