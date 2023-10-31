package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectagle Rectangle) float64 {
	return 2 * (rectagle.Width + rectagle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
