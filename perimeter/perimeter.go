package perimeter

import "math"

// // Perimeter 计算周长
// func Perimeter(rectangle Rectangle) float64 {
// 	return 2 * (rectangle.Height + rectangle.Width)
// }

// // Area 计算面积
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }

// Area 计算面积
// GO语言没有重载这一说
// func Area(circle Circle) float64 {
// 	return 2 * 3.14 * circle.Radio
// }

// Rectangle 举行对象
// 可以用rectangle.field访问
type Rectangle struct {
	Height float64
	Width  float64
}

// Area Rectangle的计算面积方法Area
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Perimeter Rectangle计算周长的方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// Circle 圆形
type Circle struct {
	Radio float64
}

// Area Circle计算面积的方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

// Perimeter Circle计算周长的方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radio
}
