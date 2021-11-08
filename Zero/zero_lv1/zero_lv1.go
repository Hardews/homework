package main

import "fmt"

func main()  {
	var inputNumber1,inpurNumber2 int
	var cmd string
	fmt.Scan(&inputNumber1,&cmd,&inpurNumber2)
	switch cmd {
	case "+":fmt.Println(inputNumber1 + inpurNumber2)
	case "-":fmt.Println(inputNumber1 - inpurNumber2)
	case "*":fmt.Println(inputNumber1 * inpurNumber2)
	case "/":fmt.Println(inputNumber1 / inpurNumber2)
	default:
		fmt.Println("error!")
	}
}
