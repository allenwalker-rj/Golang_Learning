package main

import "fmt"

func main() {
	fmt.Println("unit 3 start")

	// 数组存储实现
	var scores [5]int
	scores[0] = 95
	scores[1] = 91
	scores[2] = 39
	scores[3] = 60
	scores[4] = 21

	sum := 0
	for i := 0; i < len(scores); i++ {
		// fmt.Printf("scores[%d] = %d\n", i, scores[i])
		sum += scores[i]
	}
	avg := sum / len(scores)
	// score1 := 95
	// score2 := 91
	// score3 := 39
	// score4 := 60
	// score5 := 21

	// // 求和
	// sum := score1 + score2 + score3 + score4 + score5

	// // 平均
	// avg := sum / 5
	fmt.Printf("成绩的总和: %v, 成绩的平均值: %v", sum, avg)
}
