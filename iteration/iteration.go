package iteration

func Repeat(charter string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += charter
	}
	return repeated
}
