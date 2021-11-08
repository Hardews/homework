package main

import (
	"fmt"
)

type storage struct {
	inputNumber int
}
type Number []storage
func (n Number) Len() int{
	return len(n)
}
func (n Number) Less(i,j int) bool{
	return n[i].inputNumber > n[j].inputNumber
}
func (n Number) Swap(i,j int)  {
	n[i].inputNumber,n[j].inputNumber = n[j].inputNumber,n[i].inputNumber
}
func main()  {
	var inputNumber storage
	fmt.Println("input")
	fmt.Scan(&inputNumber)
	//不会了给我整的
}