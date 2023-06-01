package main

import "fmt"

func main() {
	// 切片 slice 对数组一个连续片段的引用
	fmt.Println("unit slice 1 start")
	_sliceTest_1()
	_sliceTest_2()
	_sliceTest_3()
	_sliceTest_4()

}

func _sliceTest_1() {
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	// 切片 是构建在数组之上
	// [1:4]切片-切出的一个片段 - 索引从1开始，到4结束，不包含4 - [1:4)
	slice := arr[1:4]
	// 输入的切片
	fmt.Println("slice:", slice)
	// 切片的长度
	fmt.Println("slice长度:", len(slice))
	// 切片的容量
	fmt.Println("slice容量:", cap(slice))
}

func _sliceTest_2() {
	// 切片并不是一个数组，而是数组的引用，具体说是一个结构体
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr下标为1的地址:%p\n", &arr[1])
	slice := arr[1:4]
	fmt.Printf("slice下标为1的地址:%p\n", &slice[0])

	// 修改切片，同时也会修改原数组的值
	slice[1] = 100
	fmt.Println("arr :", arr)
	fmt.Println("slice :", slice)
}

func _sliceTest_3() {
	// 通过 make 内置函数来创建切片，基本语法：var 切片 = make([]类型,长度,容量)
	slice := make([]int, 4, 20)
	fmt.Println(slice)

	// 赋值
	slice[0] = 11
	slice[1] = 22
	slice[2] = 33
	slice[3] = 44
	fmt.Println(slice)
	// ps：make底层创建一个数组，对外不可见，所以不可以直接操作这个数组，需要通过切片来间接访问各个元素
}

func _sliceTest_4() {
	// 定义一个切片，直接就指定具体数组，是引用原理类似make的方式
	slice := []int{1, 4, 7}
	fmt.Println(slice)
	fmt.Println("slice的长度:", len(slice))
	fmt.Println("slice的容量:", cap(slice))
}
