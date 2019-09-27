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
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0
		if got != want {
			t.Errorf("want '%.2f' but gotgit '%.2f'", want, got)
		}
	})

	t.Run("cicles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("want '%f' but got '%f'", want, got)
		}

	})

}
