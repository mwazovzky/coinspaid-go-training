package lesson1

func Generator(rule func(int) int) func() int {
	value := 1
	return func() int {
		current := value
		value = rule(value)
		return current
	}
}
