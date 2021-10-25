package str

func IndexOf(haystack string, needle string) int {
	M := len(needle)

	for i := 0; i <= len(haystack)-M; i++ {
		j := 0
		for j < M {
			if haystack[i+j] != needle[j] {
				break
			}
			j += 1
		}
		if j == M {
			return i
		}
	}
	return -1
}
