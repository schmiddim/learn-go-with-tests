package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(width, height float64) (result float64) {
	result = (width + height) * 2

	return
}

func Area(width, height float64) (result float64) {

	result = width * height
	return
}
