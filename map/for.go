package _map


func RangeForMap1(m map[int]string) {
	for k, v := range m {
		_, _ = k, v
	}

}

func RangeForMap2(m map[int]string) {
	for k, _ := range m {
		_, _ = k, m[k]
	}
}

func RangeForMap3(m map[int]string)  {
	for range m{

	}
}