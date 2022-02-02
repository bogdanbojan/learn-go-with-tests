package iteration

// Repeat returns the string rcount times.
func Repeat(s string, rcount int) string {
	var repeated string
	for i := 0; i < rcount; i++ {
		repeated += s
	}
	return repeated
}
