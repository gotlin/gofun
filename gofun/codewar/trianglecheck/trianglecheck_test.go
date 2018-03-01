package trianglecheck

import (
	"testing"
	"fmt"
)

func TestCalcPerimeter(t *testing.T) {

	res := CalcPerimeter(1, 2, 3)
	fmt.Println(res)

}

func TestCalcArea(t *testing.T) {
	res := CalcArea(3, 4, 5)
	fmt.Println(res)

}
