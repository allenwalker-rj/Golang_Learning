package main

import "fmt"

func main() {

	fmt.Println("Hello")
	var age int = 18
	if age >= 18 {
		fmt.Println("满足条件")
	} else {
		fmt.Println("不满足条件")
	}

	// 在Golang里，if后面可以并列的加入变量的定义
	if count := 20; count < 30 {
		fmt.Println("年龄不满足条件")
	}
}
