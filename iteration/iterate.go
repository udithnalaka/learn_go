package iteration

func Repeat(letter string) string {
	var value string
	for i := 0; i < 5; i++ {
		value += letter
	}

	return value

}
