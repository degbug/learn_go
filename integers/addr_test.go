package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

//ExampleAddr 示例 Example开头，要测试的方法结尾
// //Output: 6 注释是必须的，否则运行测试用例时不会执行,
// 这里会将控制台输出的结果和注释的对比，一致才通过测试
// 调用 go test -v执行测试用例
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6

}
