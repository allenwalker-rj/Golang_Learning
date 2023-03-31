package main

import (
	"fmt"
	"test-go/day04/uinit1/crm/db"
	"test-go/day04/uinit1/crm/util"
)

func main() {
	fmt.Println("import simple test")
	db.GetConnection()

	sum := util.Sum(1, 2, 3, 4, 5)
	fmt.Println(sum)
}
