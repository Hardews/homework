package main

import "fmt"

func main() {
	over := make(chan bool,1) //让它带个缓存，
	// 然后执行最后一项的时候不会堵塞
	for i := 0; i < 10; i++ {
			fmt.Println(i) //这里为啥要把那个匿名函数去掉我还没懂
		                   //就感觉它挡路了
			if i == 9 {
			over <- true
		}
	}
	<-over
	fmt.Println("over!!!")
}


