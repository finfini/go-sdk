package difflib

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateRatio(matches, length int) float64 {
	if length > 0 {
		return 2.0 * float64(matches) / float64(length)
	}
	return 1.0
}

// ListifyString convert string into slices of strings
func ListifyString(str string) (lst []string) {
	lst = make([]string, len(str))
	for i, c := range str {
		lst[i] = string(c)
	}
	return lst
}
