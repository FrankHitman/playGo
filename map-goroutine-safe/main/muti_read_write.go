package main

//func main() {
//	m := make(map[int]int)
//	go func() {
//		for {
//			_ = m[1]
//		}
//	}()
//	go func() {
//		for {
//			m[2] = 2
//		}
//	}()
//	select {}
//}


func main() {
	Map := make(map[int]int)

	for i := 0; i < 100000; i++ {
		//fmt.Println("cycle ", i)
		go writeMap(Map, i, i)
		go readMap(Map, i)
	}

}

func readMap(Map map[int]int, key int) int {
	return Map[key]
}

func writeMap(Map map[int]int, key int, value int) {
	Map[key] = value
}