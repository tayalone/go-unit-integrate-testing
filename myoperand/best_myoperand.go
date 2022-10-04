package myoperand

import (
	"github.com/tayalone/go-unit-integrate-test/star"
	"github.com/tayalone/go-unit-integrate-test/triangle"
)

/*MoItf is interface*/
type MoItf interface {
	Operand(a int, b int, c int, d int) int
}

/*MoStr is Interface*/
type MoStr struct {
	star     star.StrItf
	triangle triangle.TrgItf
}

var bs *MoStr

/*New return My Operand*/
func New(star star.StrItf, triangle triangle.TrgItf) *MoStr {
	bs = &MoStr{
		star:     star,
		triangle: star,
	}
	return bs
}

/*Operand return My Int*/
func (mo *MoStr) Operand(a int, b int, c int, d int) int {
	return mo.star.Operand(a, b) + mo.triangle.Operand(c, d)
}
