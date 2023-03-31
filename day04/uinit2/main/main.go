package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("panic test 1")
	defer test1()
	// 异常捕获后，程序依然可继续执行
	fmt.Println("panic test 2")
	err := test2()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("panic test 3")
}

// 错误的方法测试
func test1() {
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

// 自定义错误
func test2() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		// 抛出自定义错误
		return errors.New("除数不能为0")
	}
	result := num1 / num2
	fmt.Println(result)
	return nil
}
