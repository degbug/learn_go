package sums

import (
	"reflect"
	"testing"
)

//调用 go test -cover 执行覆盖率的测试检查
func TestSum(t *testing.T) {
	//两种方式指定数组
	//1.[N]type{val1,val2,val3} e.g. [3]int{1,2,3}
	//2.[...]type{val1,val2,val3} e.g. [...]int{1,2,3}
	//3.[]type[val1,val2,val3]
	// numbers := []int{1, 2, 3, 4, 5}

	// got := Sum(numbers)
	// want := 15
	// if got != want {
	// 	t.Errorf("want '%d' but got '%d'", want, got)
	// }

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15
		if got != expected {
			t.Errorf("expect '%d' but got '%d'", expected, got)
		}
	})

	t.Run("collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expected := 6
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5})
	// want := []int{3, 9}
	// if got != want {
	//DeepEqual不是类型安全的，有时会发生怪异的现象
	// want := "bob"
	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("want '%v' but got '%v' ", want, got)
	// }

	want := []int{6, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want '%v' but got '%v'", want, got)
	}
}
