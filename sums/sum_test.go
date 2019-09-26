package sums

import (
	"testing"
)

func TestSum(t *testing.T) {
	//两种方式指定数组
	//1.[N]type{val1,val2,val3} e.g. [3]int{1,2,3}
	//2.[...]type{val1,val2,val3} e.g. [...]int{1,2,3}
	//3.[]type[val1,val2,val3]
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("want '%d' but got '%d'", want, got)
	}
}
