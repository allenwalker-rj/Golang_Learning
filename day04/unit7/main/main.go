package main

import "fmt"

func main() {
	fmt.Println("unit 7 start")
	/** 长度属于类型的一部分 */
	// 定义一个数组
	// var arr1 = [3]int{3, 6, 9}
	// fmt.Printf("数组的类型为：%T", arr1)

	// var arr2 = [6]int{3, 6, 9, 1, 4, 7}
	// fmt.Printf("数组的类型为：%T", arr2)

	// /** Go中数组数组属值类型，在默认情况下是值传递，因此会进行值拷贝 */
	var arr3 = [3]int{3, 6, 9}
	test1(arr3)
	// 此时这里的test1方法并未改变原生数组的值
	fmt.Println(arr3)

	/** 如果在其他函数中，去修改原来的数组，可以使用引用传递（指针方式） */
	var arr4 = [3]int{3, 6, 9}
	// 传入arr4数组的地址
	test2(&arr4)
	fmt.Println(arr4)
}

func test1(arr [3]int) {
	arr[0] = 100
	fmt.Println(arr)
}

func test2(arr *[3]int) {
	arr[0] = 100
	fmt.Println(arr)
}
