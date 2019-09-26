package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaa"
	if repeated != expected {
		t.Errorf("expect '%s' but got '%s'", expected, repeated)
	}
}

//示例
func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	//Output: aaaa
}

//基准测试,以Benchmark开头+要测试的函数命名
//用go test -bench=.   来运行
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
