package tools

import (
	"fmt"
	"math/rand"
)

// 生成验证码
func Random4Num() string {
	intn := rand.Intn(9999)
	if intn < 1000 {
		intn += 1000
	}
	return fmt.Sprintf("%d", intn)
}
