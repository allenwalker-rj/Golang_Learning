package main

import "fmt"

func main() {
	fmt.Println("unit 8 start")
	/** 二维数组 */

	// 二维数组定义
	var arr [2][3]int16
	fmt.Println(arr)

	fmt.Printf("arr的地址是: %p\n", &arr)
	fmt.Printf("arr[0]的地址是: %p\n", &arr[0])
	fmt.Printf("arr[0][0]的地址是: %p\n", &arr[0][0])

	fmt.Printf("arr[1]的地址是: %p\n", &arr[1])
	fmt.Printf("arr[1][0]的地址是: %p\n", &arr[1][0])

	// 赋值
	arr[0][0] = 100
	arr[0][1] = 200
	arr[1][1] = 300
	fmt.Println(arr)

	// 二维数组的初始化
	var arr1 [2][3]int16 = [2][3]int16{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr1)
}
