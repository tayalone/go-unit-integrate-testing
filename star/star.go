package star

import "math"

/*Star Operation */
func Star(x int, y int) int {
	if y == 0 {
		return x * 2
	}
	if x == 0 {
		return y * 3
	}
	return int(math.Pow(float64(x/y), 3))

}
