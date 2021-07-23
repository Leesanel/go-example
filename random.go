package main
/*
  随机生成指定长度的包含数字、大写字母、小写字母的字符串
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func randString(n int) string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = s[rand.Intn(len(s))]
	}
	return string(b)
}

func main()  {
	str := randString(10)
	fmt.Println(str)
}
