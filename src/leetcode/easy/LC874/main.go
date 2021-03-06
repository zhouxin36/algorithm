package main

import "fmt"

func main() {
	fmt.Println(robotSim([]int{-2, -1, -2, 3, 7}, [][]int{{1, -3}, {2, -3}, {4, 0}, {-2, 5}, {-5, 2}, {0, 0}, {4, -4}, {-2, -5}, {-1, -2}, {0, 2}}) == 100)
	fmt.Println(robotSim([]int{1, 2, -2, 5, -1, -2, -1, 8, 3, -1, 9, 4, -2, 3, 2, 4, 3, 9, 2, -1, -1, -2, 1, 3, -2, 4, 1, 4, -1, 1, 9, -1, -2, 5, -1, 5, 5, -2, 6, 6, 7, 7, 2, 8, 9, -1, 7, 4, 6, 9, 9, 9, -1, 5, 1, 3, 3, -1, 5, 9, 7, 4, 8, -1, -2, 1, 3, 2, 9, 3, -1, -2, 8, 8, 7, 5, -2, 6, 8, 4, 6, 2, 7, 2, -1, 7, -2, 3, 3, 2, -2, 6, 9, 8, 1, -2, -1, 1, 4, 7}, [][]int{{-57, -58}, {-72, 91}, {-55, 35}, {-20, 29}, {51, 70}, {-61, 88}, {-62, 99}, {52, 17}, {-75, -32}, {91, -22}, {54, 33}, {-45, -59}, {47, -48}, {53, -98}, {-91, 83}, {81, 12}, {-34, -90}, {-79, -82}, {-15, -86}, {-24, 66}, {-35, 35}, {3, 31}, {87, 93}, {2, -19}, {87, -93}, {24, -10}, {84, -53}, {86, 87}, {-88, -18}, {-51, 89}, {96, 66}, {-77, -94}, {-39, -1}, {89, 51}, {-23, -72}, {27, 24}, {53, -80}, {52, -33}, {32, 4}, {78, -55}, {-25, 18}, {-23, 47}, {79, -5}, {-23, -22}, {14, -25}, {-11, 69}, {63, 36}, {35, -99}, {-24, 82}, {-29, -98}, {-50, -70}, {72, 95}, {80, 80}, {-68, -40}, {65, 70}, {-92, 78}, {-45, -63}, {1, 34}, {81, 50}, {14, 91}, {-77, -54}, {13, -88}, {24, 37}, {-12, 59}, {-48, -62}, {57, -22}, {-8, 85}, {48, 71}, {12, 1}, {-20, 36}, {-32, -14}, {39, 46}, {-41, 75}, {13, -23}, {98, 10}, {-88, 64}, {50, 37}, {-95, -32}, {46, -91}, {10, 79}, {-11, 43}, {-94, 98}, {79, 42}, {51, 71}, {4, -30}, {2, 74}, {4, 10}, {61, 98}, {57, 98}, {46, 43}, {-16, 72}, {53, -69}, {54, -96}, {22, 0}, {-7, 92}, {-69, 80}, {68, -73}, {-24, -92}, {-21, 82}, {32, -1}, {-6, 16}, {15, -29}, {70, -66}, {-85, 80}, {50, -3}, {6, 13}, {-30, -98}, {-30, 59}, {-67, 40}, {17, 72}, {79, 82}, {89, -100}, {2, 79}, {-95, -46}, {17, 68}, {-46, 81}, {-5, -57}, {7, 58}, {-42, 68}, {19, -95}, {-17, -76}, {81, -86}, {79, 78}, {-82, -67}, {6, 0}, {35, -16}, {98, 83}, {-81, 100}, {-11, 46}, {-21, -38}, {-30, -41}, {86, 18}, {-68, 6}, {80, 75}, {-96, -44}, {-19, 66}, {21, 84}, {-56, -64}, {39, -15}, {0, 45}, {-81, -54}, {-66, -93}, {-4, 2}, {-42, -67}, {-15, -33}, {1, -32}, {-74, -24}, {7, 18}, {-62, 84}, {19, 61}, {39, 79}, {60, -98}, {-76, 45}, {58, -98}, {33, 26}, {-74, -95}, {22, 30}, {-68, -62}, {-59, 4}, {-62, 35}, {-78, 80}, {-82, 54}, {-42, 81}, {56, -15}, {32, -19}, {34, 93}, {57, -100}, {-1, -87}, {68, -26}, {18, 86}, {-55, -19}, {-68, -99}, {-9, 47}, {24, 94}, {92, 97}, {5, 67}, {97, -71}, {63, -57}, {-52, -14}, {-86, -78}, {-17, 92}, {-61, -83}, {-84, -10}, {20, 13}, {-68, -47}, {7, 28}, {66, 89}, {-41, -17}, {-14, -46}, {-72, -91}, {4, 52}, {-17, -59}, {-85, -46}, {-94, -23}, {-48, -3}, {-64, -37}, {2, 26}, {76, 88}, {-8, -46}, {-19, -68}}) == 5140)
	fmt.Println(robotSim([]int{7, -2, -2, 7, 5}, [][]int{{-3, 2}, {-2, 1}, {0, 1}, {-2, 4}, {-1, 0}, {-2, -3}, {0, -3}, {4, 4}, {-3, 3}, {2, 2}}) == 4)
	fmt.Println(robotSim([]int{-2, 8, 3, 7, -1}, [][]int{{-4, -1}, {1, -1}, {1, 4}, {5, 0}, {4, 5}, {-2, -1}, {2, -5}, {5, 1}, {-3, -1}, {5, -3}}) == 324)
	fmt.Println(robotSim([]int{-2, -1, 8, 9, 6}, [][]int{{-1, 3}, {0, 1}, {-1, 5}, {-2, -4}, {5, 4}, {-2, -3}, {5, -1}, {1, -1}, {5, 5}, {5, 2}}) == 0)
	fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}) == 65)
	fmt.Println(robotSim([]int{4, -1, 3}, [][]int{}) == 25)
}
func robotSim(commands []int, obstacles [][]int) int {
	xm := map[int][]int{}
	ym := map[int][]int{}
	for _, ob := range obstacles {
		xm[ob[0]] = append(xm[ob[0]], ob[1])
		ym[ob[1]] = append(ym[ob[1]], ob[0])
	}
	// 0上，1右，2下，3左
	direct := 0
	pos := []int{0, 0}
	max := 0
	for _, c := range commands {
		switch {
		case c == -1:
			direct = (direct + 5) % 4
		case c == -2:
			direct = (direct + 3) % 4
		case c >= 0 && c <= 9:
			if c != 0 {
				switch direct {
				//上
				case 0:
					pos[1] = getMoves(pos[1], pos[1]+c, xm[pos[0]])
					//右
				case 1:
					pos[0] = getMoves(pos[0], pos[0]+c, ym[pos[1]])
					//下
				case 2:
					pos[1] = getMoves(pos[1], pos[1]-c, xm[pos[0]])
					//左
				case 3:
					pos[0] = getMoves(pos[0], pos[0]-c, ym[pos[1]])
				default:
				}
			}
		default:
		}
		v := pos[0]*pos[0] + pos[1]*pos[1]
		if max < v {
			max = v
		}
	}
	return max
}
func getMoves(start, end int, ob []int) int {
	if ob == nil {
		return end
	}
	for _, o := range ob {
		if start > end {
			if o >= end && o < start {
				return o + 1
			}
		} else {
			if o > start && o <= end {
				return o - 1
			}
		}
	}
	return end
}

func robotSim2(commands []int, obstacles [][]int) int {
	var offset uint = 1 << 15
	o := make(map[uint]bool, 0)
	for _, v := range obstacles {
		o[(uint(v[0])+offset)<<16|(uint(v[1])+offset)] = true
	}
	var ox, oy, max uint = offset, offset, 0
	dir := 0
	for _, v := range commands {
		switch v {
		case -2:
			dir = (dir - 1) & 0x03
		case -1:
			dir = (dir + 1) & 0x03
		default:
			switch dir {
			case 0:
				for i := 0; i < v; i++ {
					if _, ok := o[ox<<16|(oy+1)]; ok {
						break
					}
					oy++
				}
			case 1:
				for i := 0; i < v; i++ {
					if _, ok := o[(ox+1)<<16|oy]; ok {
						break
					}
					ox++
				}
			case 2:
				for i := 0; i < v; i++ {
					if _, ok := o[ox<<16|(oy-1)]; ok {
						break
					}
					oy--
				}
			case 3:
				for i := 0; i < v; i++ {
					if _, ok := o[(ox-1)<<16|oy]; ok {
						break
					}
					ox--
				}
			}
		}
		rx, ry := ox-offset, oy-offset
		r := rx*rx + ry*ry
		if r > max {
			max = r
		}
	}
	return int(max)
}
