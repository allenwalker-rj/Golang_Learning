package main

import "fmt"

func main() {
	fmt.Println("unit slice 2 start")

	slice := make([]int, 4, 20)
	slice[0] = 11
	slice[1] = 22
	slice[2] = 33
	slice[3] = 44

	// 方式1：普通for循环
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}

	fmt.Println("-------------------")

	// 方式2：range循环
	for i, v := range slice {
		fmt.Printf("slice[%v]=%v\n", i, v)
	}
}
