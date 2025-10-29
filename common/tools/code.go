package tools

import (
	"fmt"
	"math/rand"
)

// 生成4位验证码
func Gen4Code() string {
	intn := rand.Intn(9999)
	if intn < 1000 {
		intn += 1000
	}
	return fmt.Sprintf("%d", intn)
}
