package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Number struct {
	ANumber int
}
type NumberSlice []Number
func (n NumberSlice)Len()  int{
	return len(n)
}
func (n NumberSlice)Less(i,j int)  bool{
	return n[i].ANumber > n[j].ANumber
}
func (n NumberSlice)Swap(i,j int)  {
	 n[i],n[j] = n[j],n[i]
}
func main()  {
	var Number1 NumberSlice
	 rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a := Number{ANumber: rand.Intn(1000) + 1}
		Number1 = append(Number1,a)
	}
	sort.Sort(Number1)
	for _,s :=  range Number1{
		fmt.Println(s)
	}
}
