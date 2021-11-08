package main

import "fmt"

func Add(inputNumber1,inputNumber2 int,)  int{
	return inputNumber1 + inputNumber2
}
func red(inputNumber1,inputNumber2 int,)  int{
	return inputNumber1 - inputNumber2
}
func mul(inputNumber1,inputNumber2 int,)  int{
	return inputNumber1 * inputNumber2
}
func div(inputNumber1,inputNumber2 int,)  int{
	return inputNumber1 / inputNumber2
}

func main()  {
	var storage []int
	var result int
	for true {
		var inputNumber1,inputNumber2 int
		var cmd string
		fmt.Scan(&inputNumber1,&cmd,&inputNumber2)
		switch cmd {
		case "+": result = Add(inputNumber1,inputNumber2)
		case "-": result = red(inputNumber1,inputNumber2)
		case "*": result = mul(inputNumber1,inputNumber2)
		case "/": result = div(inputNumber1,inputNumber2)
		default:
			fmt.Println("error!")
		}
		storage = append(storage,result)
		fmt.Println(result)
		for _, i := range storage {
			fmt.Print(i," ")
		}
		fmt.Println()
	}

}