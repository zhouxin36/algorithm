package 字符串匹配

const PRIME int64 = 1e9 + 7
const RLength = 128

type RabinKarp struct {
	// 模式字串
	pat string
	// 模式字串hash
	patHash int64
	// 模式字串长度
	patLength int
	// 素数
	prime int64
	// 字符表长度/进制数
	R int
	//RM R进制数patLength阶乘
	RM int64
}

func NewInstanceRabinKarp(patString string) *RabinKarp {
	rk := &RabinKarp{
		pat:       patString,
		patLength: len(patString),
		prime:     PRIME,
		R:         RLength,
		RM:        0}
	var RM int64
	RM = 1
	// 最高位为长度-1
	for i := 1; i < len(patString); i++ {
		RM = (RM * int64(rk.R)) % rk.prime
	}
	rk.patHash = rk.hash(patString, rk.patLength)
	rk.RM = RM
	return rk
}
func (rk *RabinKarp) hash(patString string, length int) int64 {
	var h int64
	for i := 0; i < length; i++ {
		h = (int64(rk.R)*h + int64(patString[i])) % rk.prime
	}
	return h
}
func (rk *RabinKarp) Search(content string) int {
	contentHash := rk.hash(content, rk.patLength)
	if contentHash == rk.patHash {
		return 0
	}
	for i := rk.patLength; i < len(content); i++ {
		contentHash = (contentHash + rk.prime - (int64(content[i-rk.patLength])*rk.RM)%rk.prime) % rk.prime
		//contentHash -= int64(content[i-rk.patLength]) * rk.RM
		contentHash = (contentHash*int64(rk.R) + int64(content[i])) % rk.prime
		// 蒙特卡洛算法（概率学保证，有小概率失败）
		if contentHash == rk.patHash {
			return i - rk.patLength + 1
		}
	}
	return -1
}

func Search(context, patStr string) int {
	if len(patStr) > len(context) {
		return -1
	}
	var patHash, contextHash, RM int64
	var i int
	RM = 1
	for ; i < len(patStr); i++ {
		patHash = (patHash*RLength + int64(patStr[i])) % PRIME
		contextHash = (contextHash*RLength + int64(context[i])) % PRIME
		if i != 0 {
			// 最高位为长度-1
			RM = (RM * RLength) % PRIME
		}
	}
	if contextHash == patHash {
		return 0
	}
	for ; i < len(context); i++ {
		contextHash = (contextHash + PRIME - (int64(context[i-len(patStr)])*RM)%PRIME) % PRIME
		contextHash = (contextHash*RLength + int64(context[i])) % PRIME
		// 拉斯维加斯算法（循环匹配校验一次）
		if contextHash == patHash {
			var flag bool
			idx := i + 1 - len(patStr)
			for j := 0; j < len(patStr); j++ {
				if patStr[j] != context[idx+j] {
					flag = true
				}
			}
			if !flag {
				return i - len(patStr) + 1
			}
		}
	}
	return -1
}
