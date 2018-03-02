package trianglecheck

import "math"

func CalcPerimeter(a, b, c int) float64 {

	return float64(a + b + c)

}

func CalcArea(a, b, c int) float64 {

	s := float64(CalcPerimeter(a, b, c)) * 0.5

	res := s * (s - float64(a)) * (s - float64(b)) * (s - float64(c))

	res = math.Sqrt(res)
	return res
}

func EquableTriangle(a, b, c int) bool {
	return CalcPerimeter(a, b, c) == CalcArea(a, b, c)
}

