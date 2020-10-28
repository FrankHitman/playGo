package main

import (
	"fmt"
	"strings"

	"github.com/robfig/cron"
)

func main() {
	spec := "0 0 21 * * 0,1,2,3,4,5,6"
	d, err := cron.Parse(spec)
	if err != nil {
		fmt.Println("cron parse spec error: ", err)
	}
	fmt.Printf("schedule is %#v", d)

	// 返回的d是interface对象，需要转化为相应的struct来使用
	if dd, ok := d.(*cron.SpecSchedule); ok {
		fmt.Println("hour is ", dd.Hour)
		fmt.Println("day of month is ", dd.Dom)
		fmt.Println("day of week is ", dd.Dow)
	}

	field := "0,1,2,3,4,5,6"
	// ranges := strings.Fields(field)
	ranges := strings.FieldsFunc(field, func(r rune) bool { return r == ',' })
	fmt.Println("ranges is ", ranges)
	rangeAndStep := strings.Split(ranges[0], "/")
	lowAndHigh := strings.Split(rangeAndStep[0], "-")
	fmt.Println(rangeAndStep)
	fmt.Println(lowAndHigh)
}

// output
// schedule is &cron.SpecSchedule{Second:0x1, Minute:0x1, Hour:0x200000, Dom:0x80000000fffffffe, Month:0x8000000000001ffe, Dow:0x7f}hour is  2097152
// day of month is  9223372041149743102
// day of week is  127
// ranges is  [0 1 2 3 4 5 6]
// [0]
// [0]
