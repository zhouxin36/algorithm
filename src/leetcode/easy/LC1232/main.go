package main

import "fmt"

func main() {
	fmt.Println(!checkStraightLine([][]int{{-7, -3}, {-7, -1}, {-2, -2}, {0, -8}, {2, -2}, {5, -6}, {5, -5}, {1, 7}}))
	fmt.Println(checkStraightLine([][]int{{-3, -14}, {-4, -21}, {2, 21}, {7, 56}, {3, 28}}))
	fmt.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}))
	fmt.Println(!checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
}

type Fraction struct {
	// 分子
	Numerator int
	// 分母
	Denominator int
}

func (f *Fraction) multiply(num int) int {
	return (f.Numerator * num) / f.Denominator
}
func checkStraightLine2(coordinates [][]int) bool {
	a := Fraction{Numerator: coordinates[1][1] - coordinates[0][1], Denominator: coordinates[1][0] - coordinates[0][0]}
	if a.Denominator == 0 {
		for i := 2; i < len(coordinates); i++ {
			if coordinates[0][0] != coordinates[i][0] {
				return false
			}
		}
	} else {
		b := coordinates[0][1] - a.multiply(coordinates[0][0])
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][1] != a.multiply(coordinates[i][0])+b {
				return false
			}
		}
	}
	return true
}
func checkStraightLine(coordinates [][]int) bool {

	if len(coordinates) == 2 {
		return true
	}

	for i := 0; i < len(coordinates)-3; i++ {
		if (coordinates[i][0]-coordinates[i+1][0])*(coordinates[i+1][1]-coordinates[i+2][1]) != (coordinates[i][1]-coordinates[i+1][1])*(coordinates[i+1][0]-coordinates[i+2][0]) {
			return false
		}
	}

	return true
}
