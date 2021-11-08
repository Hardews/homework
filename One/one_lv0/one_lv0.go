package main

import "fmt"

func main()  {
	var inputNumber string
	fmt.Scan(&inputNumber)
	var bytes = []byte(inputNumber)
	for i := 0; i < len(inputNumber)/2; i++ {
		tmp := bytes[len(inputNumber) - i -1]
		bytes[len(inputNumber) - i - 1] = bytes[i]
		bytes[i] = tmp
	}
	a := string(bytes)
	fmt.Println(a)
}
