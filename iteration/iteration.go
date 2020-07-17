package iterators

const repeatAmount = 6

func Repeat(char string) string {
	var repeated string
	i := 0
	for i < repeatAmount {
		repeated += char
		i = i + 1
	}

	return repeated
}

func RepeatNTimes(char string, repeaterCount int) {
	var repeated string

	for i := 0; i < repeaterCount; i++ {
		repeated += char
	}
	return repeated
}
