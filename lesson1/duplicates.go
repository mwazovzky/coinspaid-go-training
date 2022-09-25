package lesson1

const shift = 32

func CheckDuplicates(str string) bool {
	var prev int32

	for _, ch := range str {
		diff := getDiff(ch, prev)

		if diff == 0 || diff == shift {
			return true
		}

		prev = ch
	}

	return false
}

func getDiff(a int32, b int32) int32 {
	if a > b {
		return a - b
	}

	if b > a {
		return b - a
	}

	return 0
}
