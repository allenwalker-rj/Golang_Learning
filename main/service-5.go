package main

import "fmt"

func main() {

	fmt.Println("service-5")
	// golang 里只有 for 循环，没有 while 循环
	var sum int = 0
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)

}
