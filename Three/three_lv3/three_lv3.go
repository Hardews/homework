package main

import "fmt"

func main() {
	over := make(chan bool,1)
	for i := 0; i < 10; i++ {
			fmt.Println(i)
		if i == 9 {
			over <- true
		}
	}
	<-over
	fmt.Println("over!!!")
}
