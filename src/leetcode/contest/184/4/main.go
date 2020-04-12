package main

import (
	"fmt"
)

func main() {
	fmt.Println(numOfWays(1) == 12)
	fmt.Println(numOfWays(2) == 54)
	fmt.Println(numOfWays(7) == 106494)
	fmt.Println(numOfWays(5000) == 30228214)
}

/**
Two pattern for each row 121 and 123.
We consider the state of the first row,
pattern 121: 121, 131, 212, 232, 313, 323.
pattern 123: 123, 132, 213, 231, 312, 321.
So we initialize a121 = 6, a123 = 6.

We consider the next possible for each pattern:
Patter 121 can be followed by: 212, 213, 232, 312, 313
Patter 123 can be followed by: 212, 231, 312, 232

121 => two 121, three 123
123 => two 121, two 123

So we can write this dynamic programming transform equation:
b121 = a121 * 3 + a123 * 2
b123 = a121 * 2 + a123 * 2

We calculate the result iteratively
and finally return the sum of both pattern a121 + a123


Complexity
Time O(N)
Space O(1)
*/
func numOfWays(n int) int {
	MOD := int64(1e9 + 7)
	color3 := int64(6)
	color2 := int64(6)

	for i := 2; i <= n; i++ {
		tmp := color3
		color3 = (2*color3 + 2*color2) % MOD
		color2 = (2*tmp + 3*color2) % MOD
	}
	return int((color3 + color2) % MOD)
}
