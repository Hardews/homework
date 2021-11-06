package main

import "fmt"

func Receiver (v...interface{}){ //定义一个名为Receiver的函数，
	// v...是可变参数的意思，就是可以输入任意多个变量，
	//定义了一个空接口，就是可以传入任意类型的参数
	for _,x := range v{  //遍历这些传入的参数
		switch x.(type) {  //switch判断这些函数所属类型
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		case float64:
			fmt.Println("float64")
		case float32:
			fmt.Println("float32")
		}
	}
}
func main()  {
	a := "你好吗"
	b := 123
	c := true
	fmt.Println("这个参数",a,"的类型是")
	Receiver(a)
	fmt.Println("这个参数",b,"的类型是")
	Receiver(b)
	fmt.Println("这个参数",c,"的类型是")
	Receiver(c)
}

