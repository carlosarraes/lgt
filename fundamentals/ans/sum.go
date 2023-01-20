package ans

func Sum(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}

	return sum
}

func SumAll(n ...[]int) []int {
	var sums []int
	for _, v := range n {
		sums = append(sums, Sum(v))
	}

	return sums
}

func SumAllTails(n ...[]int) []int {
	var sums []int
	for _, v := range n {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(v[1:]))
		}
	}

	return sums
}
