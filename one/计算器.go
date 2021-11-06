package main

import "fmt"

func add(Number1,Number2 int,)  int {
	return Number1 + Number2
}
func minus(Number1,Number2 int,)  int {
	return Number1 - Number2
}
func mul(Number1,Number2 int,)  int {
	return Number1 * Number2
}
func div(Number1,Number2 int,)  int {
	return Number1 / Number2
}
func main()  {
	var history []int
	for true {
		var inputNumber1, inputNumber2 int
		var cmd string
		var ans int
		fmt.Scan(&inputNumber1, &cmd, &inputNumber2)

		switch cmd {
		case "+":
			ans = add(inputNumber1, inputNumber2)
		case "-":
			ans = minus(inputNumber1, inputNumber2)
		case "*":
			ans = mul(inputNumber1, inputNumber2)
		case "/":
			ans = div(inputNumber1, inputNumber2)
		default:
			fmt.Println("error!")
		}
		history = append(history,ans)
		fmt.Println(ans)
		fmt.Println("历史结果为：")
		fmt.Println(history)
	}
}
