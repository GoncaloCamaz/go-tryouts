package iteration

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}

	return repeated
}

func RepeatNTimes(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}

	return repeated
}
