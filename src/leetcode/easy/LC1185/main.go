package main

import "fmt"

func main() {
	fmt.Println(dayOfTheWeek(29, 2, 2016) == "Monday")
	fmt.Println(dayOfTheWeek(31, 8, 2019) == "Saturday")
	fmt.Println(dayOfTheWeek(18, 7, 1999) == "Sunday")
	fmt.Println(dayOfTheWeek(15, 8, 1993) == "Sunday")
}
func dayOfTheWeek(day int, month int, year int) string {
	if month < 3 {
		month += 12
		year -= 1
	}
	week := (year + year/4 - year/100 + year/400 + 13*(month+1)/5 + day - 1) % 7
	switch week {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	default:
		return ""
	}
}
