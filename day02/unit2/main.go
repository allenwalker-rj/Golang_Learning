package main

import (
	"fmt"
	"log"
)

func main() {

	log.Println("service-6")
	num1 := 100
	num2 := 25
	var sum = sum(num1, num2)
	fmt.Println(sum)
	var substract = substract(num1, num2)
	fmt.Println(substract)
	var multiply, divide, mod = multiplyAndDivide(num1, num2)
	fmt.Println(multiply, divide, mod)
	// 如果有返回值不想接收，可以用 _ 进行忽略
	sum2, _ := sumAndSubstract(num1, num2)
	fmt.Println(sum2)
}

// 函数命名要求
// 首字母不能是数字
// 首字母小写只能被本包文件调用，其他包文件不能调用（类似 private）
// 首字母大写可以被本包文件和其他文件使用（类似 public）
func sum(n1, n2 int) int {
	return n1 + n2
}

func substract(n1, n2 int) int {
	return n1 - n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}

func divide(n1, n2 int) int {
	return n1 / n2
}

func mod(n1, n2 int) int {
	return n1 % n2
}

func sumAndSubstract(n1, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}

func multiplyAndDivide(n1, n2 int) (int, int, int) {
	return n1 * n2, n1 / n2, n1 % n2
}
