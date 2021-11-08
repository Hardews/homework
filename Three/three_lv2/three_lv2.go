package main

import (
	"fmt"
	"os"
)


func main()  {
	 file,err := os.Create("D:/GOProjects/src/Three/three_lv2/plan.exe")
	 if err!= nil{
		 fmt.Println("err:",err)
		 return
	 }

	 file.WriteString("I’m not afraid of difficulties and insist on learning programming")
	 fileName := "D:/GOProjects/src/Three/three_lv2/plan.exe"
	 file,err := os.OpenFile(fileName)
	 if err!=nil {
		fmt.Println("err:",err)
		return
	 }
	 file.Read()

}
