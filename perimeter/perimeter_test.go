package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}

		got := rectangle.Perimeter()
		want := 40.0
		if got != want {
			//f对应float64，.2表示输出两位小数
			t.Errorf("want '%.2f' but got '%.2f' ", want, got)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 62.83185307179586
		if got != want {
			//f对应float64，.2表示输出两位小数
			t.Errorf("want '%f' but got '%f' ", want, got)
		}
	})

}

func TestArea(t *testing.T) {

	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()

	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("want '%.2f' but gotgit '%.2f'", want, got)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("cicles", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	checkArea(t, circle, 314.1592653589793)

	// })

	// 表格驱动测试

	// 这里定义了一个匿名的结构体
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("want '%.2f' but gotgit '%.2f'", tt.want, got)
		}
	}
}
