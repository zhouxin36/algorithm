package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30") == 1)
	fmt.Println(daysBetweenDates("2020-01-15", "2019-12-31") == 15)
}
func daysBetweenDates(date1 string, date2 string) int {
	t1, _ := time.Parse("2006-01-02", date1)
	t2, _ := time.Parse("2006-01-02", date2)
	date := t1.Unix() - t2.Unix()
	res := int(date / 86400)
	if res > 0 {
		return res
	} else {
		return -res
	}
}
