package main

import (
	"fmt"
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

func main() {
	go sayHello()
	for i := 0; i < 5; i++ {
		fmt.Println("a")
		time.Sleep(1 * time.Second)
	}
}
