package main

//主函数所在包必须命名为main
//引入fmt包
import "fmt"

//常量
const helloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const chinese = "Chinese"
const chineseHelloPrefix = "你好, "

//主函数
func main() {
	fmt.Println(Hello("World", ""))
}

// Hello 定义函数
// 返回字符串
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return helloPrefix + name
}
