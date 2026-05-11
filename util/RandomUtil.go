package util

import "math/rand"

//生成随机数

func GetRandomNum() string {
	var str string

	for i := 0; i < 4; i++ {
		current := rand.Intn(10)

		str += String(current)
	}
	return str
}
