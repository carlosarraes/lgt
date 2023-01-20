package iteration

func Repeat(c string, n int) string {
	var repeated string

	if n <= 0 {
		return c
	}

	for i := 0; i < n; i++ {
		repeated += c
	}
	return repeated
}
