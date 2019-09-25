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

	return greetingPrefix(language) + name
}

// greetingPrefix 根据语言获取前缀
// 返回值（prefix string)可以在方法定义上定义返回值，这样可以定义多个返回值如(a,b string)
func greetingPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chineseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}
