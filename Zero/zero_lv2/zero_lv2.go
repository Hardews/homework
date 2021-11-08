package main

import "fmt"

func main()  {
	var storage  map[string]string
	storage = make(map[string]string)
	var middle map[string]string
	middle = make(map[string]string)

	storage["abcd"] = "ssss"
	storage["1234"] = "sdad"
	middle["ssss"]  =  "No.1"
	middle["sdad"]  =  "No.2"

	var cipher string
	var account string
	fmt.Println("请输入账号")
	fmt.Scan(&account)
	fmt.Println("请输入密码")
	fmt.Scan(&cipher)

	 _,ok1 := middle[account]
	if ok1 {
		_,ok2 := storage[cipher]
		if ok2 {
			fmt.Println("登陆成功")
		}else {
			fmt.Println("账号或密码错误")
		}
	}else {
		fmt.Println("账号或密码错误")
	}


}
