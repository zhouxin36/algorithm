package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09") == 9)
	fmt.Println(dayOfYear("2019-02-10") == 41)
	fmt.Println(dayOfYear("2003-03-01") == 60)
	fmt.Println(dayOfYear("2004-03-01") == 61)
}
func dayOfYear(date string) int {
	t1, _ := time.Parse("2006-01-02", date)
	t2, _ := time.Parse("2006-01-02", date[:4]+"-01-01")
	return t1.YearDay() - t2.YearDay() + 1
}
