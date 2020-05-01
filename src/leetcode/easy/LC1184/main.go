package main

import "fmt"

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3) == 4)
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1) == 1)
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2) == 3)
}
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	p, q := start, (start-1+len(distance))%len(distance)
	var pNum, qNum int
	for {
		if p != destination && q != (destination-1+len(distance))%len(distance) {
			pNum += distance[p]
			qNum += distance[q]
			p = (p + 1) % len(distance)
			q = (q - 1 + len(distance)) % len(distance)
		} else if p != destination && qNum > pNum {
			pNum += distance[p]
			p = (p + 1) % len(distance)
		} else if q != (destination-1+len(distance))%len(distance) && qNum < pNum {
			qNum += distance[q]
			q = (q - 1 + len(distance)) % len(distance)
		} else {
			break
		}

	}
	if pNum > qNum {
		return qNum
	} else {
		return pNum
	}
}
