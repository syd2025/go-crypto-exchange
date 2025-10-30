package tools

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func Md5HexString(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}
func Md5Base64String(s string) string {
	h := md5.Sum([]byte(s))
	return base64.StdEncoding.EncodeToString(h[:])
}

// GenerateSalt 生成随机盐
func generateSalt(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// HashPasswordWithSalt 对密码进行加盐MD5加密
func HashPasswordWithSalt(password string) (salt string, hash string, err error) {
	salt, err = generateSalt(8) // 8字节盐值
	if err != nil {
		return "", "", err
	}
	h := md5.New()
	io.WriteString(h, password+salt)
	hash = hex.EncodeToString(h.Sum(nil))
	return salt, hash, nil
}
