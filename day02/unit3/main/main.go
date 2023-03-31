package main

import "fmt"

// 内存分析问题
func main() {
	fmt.Println("service-7")
	num1 := 100
	num2 := 25
	fmt.Printf("交换前的两个数: num1 = %v, num2 = %v\n", num1, num2)
	exchangeNum(num1, num2)
	fmt.Printf("交换后的两个数: num1 = %v, num2 = %v\n", num1, num2)
}

func exchangeNum(n1, n2 int) {
	var t int
	t = n1
	n1 = n2
	n1 = t
}
