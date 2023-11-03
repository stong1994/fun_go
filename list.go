package gofun

func Reverse[T any](list []T) []T {
	l2 := make([]T, len(list))
	copy(l2, list)
	for i, j := 0, len(l2)-1; i < j; i, j = i+1, j-1 {
		l2[i], l2[j] = l2[j], l2[i]
	}
	return l2
}

func ReverseString(str string) string {
	runes := []rune(str)
	return string(Reverse(runes))
}
