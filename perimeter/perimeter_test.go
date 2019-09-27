package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	if got != want {
		//f对应float64，.2表示输出两位小数
		t.Errorf("want '%.2f' but got '%.2f' ", want, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0
	if got != want {
		t.Errorf("want '%.2f' but got '%.2f'", want, got)
	}
}
