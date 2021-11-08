package main

import (
	"fmt"
	"sync"
)

var x int64
var a = make(chan bool,1)
var wg sync.WaitGroup

func add() {
	for i := 0; i < 50000; i++ {
        a <- true
		x = x + 1
		<- a
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
