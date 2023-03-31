package main

import "fmt"

func main() {
	fmt.Println("unit 6 start")
	// 数组初始化方式演示
	// 第一种
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// 第二种
	var arr1 = [...]int{12, 13, 14, 15}
	fmt.Println(arr1)

	// 第三种
	var arr2 = [...]int{2: 55, 3: 66, 4: 77}
	fmt.Println(arr2)

}
