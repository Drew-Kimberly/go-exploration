package str

// IndexOf returns the starting index of the provided "needle" string within the "haystack" string.
// The index of the first occurrence of the needle will be returned.
// If the needle cannot be found in its entirety, -1 will be returned.
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
