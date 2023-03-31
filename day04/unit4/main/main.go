package main

import "fmt"

func main() {
	fmt.Println("unit 4 start")

	var arr [3]int16
	// 获取数组的长度
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	// 打印数组
	fmt.Println(arr)
	// 证明 arr 中存储的是地址值
	fmt.Printf("arr的地址为: %p\n", &arr)
	// 打印数组的第一个元素地址
	fmt.Printf("arr的地址为: %p\n", &arr[0])
	fmt.Printf("arr的地址为: %p\n", &arr[1])
	fmt.Printf("arr的地址为: %p\n", &arr[2])

}
