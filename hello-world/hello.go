package main

//主函数所在包必须命名为main
//引入fmt包
import "fmt"

//主函数
func main() {
	fmt.Println(Hello("World"))
}

// Hello 定义函数
// 返回字符串
func Hello(name string) string {
	return "Hello, " + name
}
