package utils

import "fmt"

func add(x, y int) int {
	return x + y
}

func Add(x, y int) int {
	return x + y
}

func SwapVar(a, b *int) {
	*a, *b = *b, *a
}

func CheckInt(v int) string {
	switch {
	case v > 0:
		return fmt.Sprintf("Число %d положительное", v)
	case v == 0:
		return "Число является нулем"
	default:
		return fmt.Sprintf("Число %d отрицательное", v)
	}
}

func ChangeBytes(s string) string {
	b := []byte(s)

	for i := range b {
		b[i]++
	}

	return string(b)
}

func ChangeSymbol(old, new rune, s string) string {
	r := []rune(s)

	for i := range r {
		if r[i] == old {
			r[i] = new
		}
	}

	return string(r)
}
