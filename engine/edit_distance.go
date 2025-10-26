package engine

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

// EditDistance: use levenshtein distance to calculate the distance between two strings
func EditDistance(a, b string) int {
	m, n := len(a), len(b)
	if m < n {
		// ensure that b is the smallest string
		a, b = b, a
		m, n = n, m
	}

	prev := make([]int, n+1)
	curr := make([]int, n+1)

	for j := 0; j <= n; j++ {
		prev[j] = j
	}

	for i := 1; i <= m; i++ {
		curr[0] = i
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				curr[j] = prev[j-1]
			} else {
				curr[j] = 1 + min(prev[j], curr[j-1], prev[j-1])
			}
		}
		copy(prev, curr)
	}

	return prev[n]
}
