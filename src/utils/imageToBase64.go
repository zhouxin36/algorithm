package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		//"https://pic.leetcode-cn.com/ff57274ee622baefedb96786aeb4b7e1eeae70161a2ec114297dd2b7313b211f",
		//"https://pic.leetcode-cn.com/bdf1e39068ce1f39d11d76e426600da9b86517cf6cba36812802e5c6db9f0e02",
		"https://pic.leetcode-cn.com/16bf7acb8a8687fdae8396380d9c635f48e02ef2483e99afc60acc0a7b610d72",
	}
	fmt.Println("------------------------------------------------------------------------------------------------------")
	for _, url := range urls {
		fmt.Println(url)
		fmt.Println(imageToBase64(url))
		fmt.Println("------------------------------------------------------------------------------------------------------")
	}

}

func imageToBase64(url string) string {
	//读原图片
	ff, _ := http.Get(url)
	defer ff.Body.Close()

	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Body.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	return sourcestring
	//写入临时文件
	//ioutil.WriteFile("a.png.txt", []byte(sourcestring), 0667)
	//读取临时文件
	//cc, _ := ioutil.ReadFile("a.png.txt")
	//解压
	//dist, _ := base64.StdEncoding.DecodeString(string(cc))
	//写入新文件
	//f, _ := os.OpenFile("xx.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	//defer f.Close()
	//f.Write(dist)
}
