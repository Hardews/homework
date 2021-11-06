package main

import (
	"fmt"
)

var chan1 = make(chan bool, 1)
var chan2 = make(chan bool)
var index = make(chan bool, 1) //定义三个管道，用来当锁
func dualNumber()  {
	for i := 0; i <= 50; i++ {
		<-chan1  //3.取出chan1里的值，执行下面代码
		v := 2*i
		fmt.Println(v) //4.输出一个偶数0
		chan2 <- true  //5.往chan2里存一个值，再阻塞
	}

}
func singularNumber()  {
	for i := 0; i < 50; i++ {
		<-chan2  //7.又取出，再次进行
		x := 2*i + 1
		fmt.Println(x) //8.输出奇数
		chan1 <- true  //9.再往chan1存入一个值，再阻塞，直到循环结束
	}
	index <- true  //10.往index存入一个值，结束这个循环
}
func main()  {
	chan1 <- true //1.先往chan1这个管道放一个值，
	// 阻塞运行，直到取出
	go singularNumber() //2.执行这个goroutine
	go dualNumber()  //6.当chan2阻塞，执行这里
	<-index  //11.循环结束，取出值不用，代码结束
}