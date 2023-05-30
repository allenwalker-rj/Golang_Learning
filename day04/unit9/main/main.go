package main

import "fmt"

func main() {
	/** 二维数组的遍历*/
	fmt.Println("unit 9 start")

	var arr [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(arr)

	_test1(arr)

	_test2(arr)
}

// 普通for循环
func _test1(arr [3][3]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
}

func _test2(arr [3][3]int) {
	for key, value := range arr {
		for k, v := range value {
			fmt.Printf("arr[%v][%v]=%v\n", key, k, v)
		}
		fmt.Println()
	}
}
