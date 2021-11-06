package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var s []int
	var n int
	for i := 0 ; i <= 100; i++ {
		rand.Seed(time.Now().UnixNano())
		n = rand.Intn(100) + 1
		s = append(s , n)
    }
	for i := 0; i < len(s); i++ {
		for j := 1; j < len(s)-i; j++ {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
	fmt.Println(s)
}