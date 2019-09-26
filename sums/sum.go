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

//SumAll 用...表示可变参数
func SumAll(numbersToSum ...[]int) (sums []int) {

	//方法1：
	/*
		lengthOfNumbers := len(numbersToSum)
		//make可以在创建切片是指定长度和容量
		sum = make([]int, lengthOfNumbers)
		for i, numbers := range numbersToSum {
			sum[i] = Sum(numbers)
		}
	*/

	//方法2：

	//使用append往数组中添加计算后的总和
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
