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
	res := CalcArea(1, 2, 3)
	fmt.Println(res)

}
