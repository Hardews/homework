package main

import (
	"fmt"
	"time"
)

var x int
var chan1 = make(chan bool,1)
var chan2 = make(chan bool,1)
var index = make(chan bool,1)
func main()  {
	chan1 <- true
	go SingularNumber()
	go dualNumber()
	<-index
	fmt.Println("over!")
}
func SingularNumber()  {
	for i := 0; i < 50; i++ {
		<- chan1
		time.Sleep(1*time.Second)
		x = 2 * i
		fmt.Println(x)
		chan2 <- true
	}
}
func dualNumber()  {
	for i := 0; i < 50; i++ {
		<- chan2
		time.Sleep(1*time.Second)
		x = 2 * i + 1
		fmt.Println(x)
		chan1 <- true
	}
	index <- true
}
