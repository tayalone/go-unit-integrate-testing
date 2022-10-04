package star

import "math"

/*StrItf is Interface*/
type StrItf interface {
	Operand(x int, y int) int
}

/*BestStr is Interface*/
type BestStr struct{}

var bs *BestStr

/*New retun Pointer of  BestStr*/
func New() *BestStr {
	bs = &BestStr{}
	return bs
}

/*Operand retun int*/
func (bs *BestStr) Operand(x int, y int) int {
	if y == 0 {
		return x * 2
	}
	if x == 0 {
		return y * 3
	}
	return int(math.Pow(float64(x/y), 3))
}
