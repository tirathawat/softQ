package sum

func Sum(xs ...int) int {
	var total int
	for _, x := range xs {
		total += x
	}
	return total
}
