package main

import "fmt"

func main() {
	fmt.Println("unit slice  3 start")

	// 定义数组
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	// 定义切片
	var slice []int = arr[1:4]
	fmt.Println(slice)

	// 追加元素的时候对数组进行扩容，旧数组扩容为新数组
	// 创建一个新的数组，将旧数组中的 1,2,3 复制到新数组中，在新数组中追加元素
	// slice2 底层数组的指向是一个新的数组
	slice2 := append(slice, 50, 100)
	fmt.Println(slice2)

	// 如果需要修改原数组，可使用以下方式
	slice = append(slice, 33, 55)
	fmt.Println(slice)
	fmt.Println(arr)

	copyTest()

}

func copyTest() {
	fmt.Println("copyTest start")
	var a []int = []int{1, 2, 3, 4, 5, 6}
	var b []int = make([]int, 10)

	// 拷贝
	copy(b, a)
	fmt.Println(b)
}
