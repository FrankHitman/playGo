package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("     len cap   address")
	fmt.Print("111---", len(nums), cap(nums))
	fmt.Printf("    %p\n", nums) //0xc4200181e0
	a := nums[:3]
	fmt.Print("222---", len(a), cap(a))
	fmt.Printf("    %p\n", a) //0xc4200181e0 一样
	//output: 222--- 3 5

	//b := nums[:3:3]          //第二个冒号 设置cap的
	var b = make([]int, len(nums[:3:3]))
	n := copy(b, nums[:3:3]) //第二个冒号 设置cap的
	fmt.Print("333---", len(b), cap(b))
	fmt.Printf("    %p\n", b) //0xc4200181e0 一样
	//output: 333--- 3 3
	fmt.Println(n, b)
	nums[0] = 55
	fmt.Println(nums, a, b)
}

//------b := nums[:3:3]  output-----
//len cap   address
//111---5 5    0xc000080000
//222---3 5    0xc000080000
//333---3 3    0xc000080000
//3 [1 2 3]
//[55 2 3 4 5] [55 2 3] [55 2 3]

// ----var b = make([]int, len(nums[:3:3]))   output----
//len cap   address
//111---5 5    0xc000014180
//222---3 5    0xc000014180
//333---3 3    0xc0000162e0
//3 [1 2 3]
//[55 2 3 4 5] [55 2 3] [1 2 3]
