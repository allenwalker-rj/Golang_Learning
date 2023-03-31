package db

import "fmt"

// 首字母大写，才可以被其他包访问
func GetConnection() {
	fmt.Println("get connection")
}
