package interation

const repeateCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeateCount; i++ {
		repeated = repeated + character
	}

	return repeated
}
