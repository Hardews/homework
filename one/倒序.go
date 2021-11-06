package main

import "fmt"

func main()  {
	var letter string
	fmt.Scan(&letter)
	var bytes = []byte(letter)
	for i := 0 ; i < len(letter)/2; i++ {
		tmp := bytes[len(letter)-i-1]
		bytes[len(letter)-i-1] = bytes[i]
		bytes[i] = tmp
	}
	letter = string(bytes)
	fmt.Println("您想要的结果是：",letter)
}
