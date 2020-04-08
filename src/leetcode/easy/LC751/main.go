package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//[]string{"255.0.0.7/32","255.0.0.8/29","255.0.0.16/32"}
	fmt.Println(ipToCIDR("255.0.0.7", 10))
}

func ipToCIDR(ip string, n int) []string {
	var result []string
	var ips []string = strings.Split(ip, ".")
	var x int64
	for _, s := range ips {
		num, _ := strconv.Atoi(s)
		x = int64(num) + x*256
	}
	for n > 0 {
		step := int(x & -x)
		for step > n {
			step >>= 1
		}
		result = append(result, longToStr(x, step))
		x += int64(step)
		n -= step
	}
	return result
}

func longToStr(x int64, step int) string {
	var ans [4]int
	//&255操作求出后8位十进制表示
	ans[0] = (int)(x & 255)
	//右移8位，即求下一个块
	x >>= 8
	ans[1] = (int)(x & 255)
	x >>= 8
	ans[2] = (int)(x & 255)
	x >>= 8
	ans[3] = int(x)
	l := 33
	//每一位就可以表示2个
	for step > 0 {
		l--
		step >>= 1
	}
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(ans[3]))
	sb.WriteByte('.')
	sb.WriteString(strconv.Itoa(ans[2]))
	sb.WriteByte('.')
	sb.WriteString(strconv.Itoa(ans[1]))
	sb.WriteByte('.')
	sb.WriteString(strconv.Itoa(ans[0]))
	sb.WriteByte('/')
	sb.WriteString(strconv.Itoa(l))
	return sb.String()
}
