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
	if age2 := 45; age2 <= 35 {
		fmt.Println("年龄满足条件")
	} else {
		fmt.Println("年龄不满足条件")
	}

	if count := 80; count >= 100 {
		fmt.Println("你的成绩等级为: S")
	} else if count >= 90 {
		fmt.Println("你的成绩等级为: A")
	} else if count >= 80 {
		fmt.Println("你的成绩等级为: B")
	} else if count >= 70 {
		fmt.Println("你的成绩等级为: C")
	} else if count >= 60 {
		fmt.Println("你的成绩等级为: D")
	} else {
		fmt.Println("你的成绩等级为: 天才")
	}
}
