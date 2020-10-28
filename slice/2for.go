package slice

func ForSlice(s []string) {
	length := len(s)
	for i := 0; i < length; i++ {
		_, _ = i, s[i]
	}
}

func RangeForSlice(s []string) {
	for i, v := range s {
		_, _ = i, v
	}
}

func RangeForSlice2(s []string) {
	for i, _ := range s {
		_, _ = i, s[i]
	}
}
