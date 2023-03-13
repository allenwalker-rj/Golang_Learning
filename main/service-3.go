package main

import "fmt"

func main() {

	// & : 返回变量的存储地址
	// * : 返回指针变量对应的数值
	fmt.Println("aaaa")

	var age int = 18
	fmt.Println("age 对应的存储空间地址为：", &age)

	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr 这个指针指向的具体数值为：", *ptr)
}
