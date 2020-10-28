package main

import "fmt"

// 遍历slice时经常用到range操作，range会复制range的对象。
// 下面例子中在循环内部改变slice的属性，最终会作用到slice上导致最后输出[1 2 101]。
// 但是并不会导致循环在第三次就结束，因为range s是从s的复本中读取i和n的。
// s的复本只复制了指针，底层元素仍指向同一片，因此可以在循环内改变slice元素的值并在循环期内可见。

func main() {
	s := []int{1, 2, 3, 4, 5}
	for i, n := range s {
		if i == 0 {
			s = s[:3]
			s[2] = n + 100
		}
		fmt.Println(i, n) // 输出1 2;2 101;3 4;4 5
	}
	fmt.Println(s) //输出 1 2 101
}


// ----output----
//0 1
//1 2
//2 101
//3 4
//4 5
//[1 2 101]
