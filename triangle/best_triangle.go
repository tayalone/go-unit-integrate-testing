package triangle

/*TrgItf is Interface*/
type TrgItf interface {
	Operand(x int, y int) int
}

/*BestTrg is Interface*/
type BestTrg struct{}

var bt *BestTrg

/*New retun Pointer of  BestTrg*/
func New() *BestTrg {
	bt = &BestTrg{}
	return bt
}

/*Operand retun int*/
func (bs *BestTrg) Operand(x int, y int) int {
	if y == 0 {
		return 1
	}
	if x == 0 {
		return 3
	}
	return (x * y) / 2
}
