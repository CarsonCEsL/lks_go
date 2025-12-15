package main

import (
	"fmt"
	"sync"
	"time"
)

var s1 string = `test
test1
test2`

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello world")
		time.Sleep(1 * time.Second)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	c <- sum + 10
}

func test(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func test2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println("c", x)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d start.\n", id)
	fmt.Printf("worker %d finish.\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() //等待所有的goroutine完成
	fmt.Println("all done")
}
