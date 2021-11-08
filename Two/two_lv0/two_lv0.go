package main

import "fmt"

type star struct {
	name     string
	diameter float64
	age      int
}

func main()  {
	earth := star{
		name: "earth",
		diameter: 12742,
		age: 20000,
	}
	fmt.Println(earth)
	earth.age +=1
	fmt.Println(earth.age)

}