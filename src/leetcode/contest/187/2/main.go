package main

func main() {

}
func kLengthApart(nums []int, k int) bool {
	var count, i int
	for ; i < len(nums); i++ {
		if nums[i] == 1 {
			break
		}
	}
	i++
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			if count < k {
				return false
			}
			count = 0
		}
	}
	return true
}
