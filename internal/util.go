package internal

import "math"

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func SplitStringByN(s string, n int) []string {
	if n <= 0 {
		return []string{} // Handle invalid n
	}
	if s == "" {
		return []string{""} // Handle empty string
	}

	var chunks []string

	for i := 0; i < len(s); i += n {
		end := i + n
		chunks = append(chunks, s[i:end])
	}
	return chunks
}