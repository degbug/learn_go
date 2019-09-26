package sums

// Sum add
// func Sum(numbers []int) int {
// 	var sum int
// 	for i := 0; i < len(numbers); i++ {
// 		sum += numbers[i]
// 	}

// 	return sum
// }

//第1次重构

//Sum 实现数组相加
//
func Sum(numbers []int) int {
	sum := 0
	// for index, number := range numbers {
	// 	sum += number
	// 	fmt.Println(index)
	// }
	//用range简化，range每次会返回数组的下标和数组元素值，用_ 空白标志符号来忽略索引
	for _, number := range numbers {
		sum += number
	}

	return sum
}
