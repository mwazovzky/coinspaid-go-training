package lesson1

func Next() func() int {
	a := 1
	b := 1

	return func() (next int) {
		next = a
		a = b
		b = b + next
		return
	}
}
