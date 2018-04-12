package iteration

// Repeat character n times
func Repeat(str string, n int) (repeated string) {
	for i := 0; i < n; i++ {
		repeated += str
	}
	return
}
