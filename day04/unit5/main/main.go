package main

import "fmt"

func main() {
	fmt.Println("unit 5 start")

	var scores [5]int
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请录入第 %d 个学生成绩: ", i+1)
		_, err := fmt.Scanln(&scores[i])
		if err != nil {
			return
		}
	}

	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	avg := sum / len(scores)
	fmt.Printf("成绩的总和: %v, 成绩的平均值: %v", sum, avg)

	fmt.Println("---------------------------")
	for index, value := range scores {
		fmt.Printf("第 %d 个学生的成绩为: %d\n", index+1, value)
	}
}
