package str

func Reverse(s string) string {
	a := []rune(s)
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
	return string(a)
}
