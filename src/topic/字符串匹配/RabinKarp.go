package 字符串匹配

const PRIME = 1e9 + 7
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
		if contentHash == rk.patHash {
			return i - rk.patLength + 1
		}
	}
	return -1
}
