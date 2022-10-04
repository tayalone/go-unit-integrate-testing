package myoperand

import (
	"github.com/tayalone/go-unit-integrate-test/star"
	"github.com/tayalone/go-unit-integrate-test/triangle"
)

/*MyOperand retrurn int*/
func MyOperand(a int, b int, c int, d int) int {
	return star.Star(a, b) + triangle.Triagle(c, d)
}
