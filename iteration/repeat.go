package iteration

// Repeat repeat four time
func Repeat(character string) string {

	// repeated := ""
	//定义变量
	var repeated string
	//循环，go只有for，没有while
	for i := 0; i < 4; i++ {
		// repeated = repeated + character
		// += 自加操作
		repeated += character
	}

	return repeated
}
