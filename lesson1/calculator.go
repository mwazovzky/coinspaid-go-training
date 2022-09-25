package lesson1

import (
	"strconv"
)

/**
 * @todo:
 * обработка ошибок - следить за количеством и последовательностью параметров, Eval(... any)
 * определять тип возвращаемого значения в зависимости от типа аргументов
 */
func Eval(v any, params ...any) any {
	var result float64
	var value float64

	result = getFloat64(v)

	for _, param := range params {
		if res, ok := param.(func(float64, float64) float64); ok {
			result = res(result, value)
			continue
		}

		value = getFloat64(param)
	}

	return result
}

func Plus(a float64, b float64) float64 {
	return a + b
}

func Minus(a float64, b float64) float64 {
	return a - b
}

func getFloat64(value any) float64 {
	if v, ok := value.(float64); ok {
		return v
	}

	if v, ok := value.(float32); ok {
		return float64(v)
	}

	if v, ok := value.(int64); ok {
		return float64(v)
	}

	if v, ok := value.(int32); ok {
		return float64(v)
	}

	if v, ok := value.(int); ok {
		return float64(v)
	}

	if v, ok := value.(string); ok {
		value, _ := strconv.ParseFloat(v, 64)
		return float64(value)
	}

	panic("Unexpected type")
}
