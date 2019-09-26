package iteration

// Repeat repeat four time
func Repeat(character string, repeatCount int) string {

	// repeated := ""
	//定义变量
	var repeated string
	//循环，go只有for，没有while
	for i := 0; i < repeatCount; i++ {
		// repeated = repeated + character
		// += 自加操作
		repeated += character
	}

	return repeated
}
