package lesson1

const shift = 32

func CheckDuplicates(str string) bool {
	var mask int32

	for _, ch := range str {
		code := getCode(ch)

		if (mask & code) != 0 {
			return true
		}

		mask = mask | code
	}

	return false
}

func getCode(ch int32) int32 {
	return 1 << getShift(ch)
}

func getShift(ch int32) int32 {
	if 'A' <= ch && ch <= 'Z' {
		return ch - 65
	}

	if 'a' <= ch && ch <= 'z' {
		return ch - 97
	}

	panic("Wrong character?")
}
