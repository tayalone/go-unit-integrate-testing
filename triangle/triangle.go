package triangle

/*Triagle return triahle area*/
func Triagle(x int, y int) int {
	if y == 0 {
		return 1
	}
	if x == 0 {
		return 3
	}

	return (x * y) / 2
}
