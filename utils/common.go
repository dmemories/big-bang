package utils

func Ternary[T any](condition bool, arg1, arg2 T) T {
	if condition {
		return arg1
	}

	return arg2
}
