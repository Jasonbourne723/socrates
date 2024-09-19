package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// generateRandomString 生成指定长度的随机字符串
// 使用base64编码来增加可读性，但请注意这会增加字符串的长度
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length*3/4) // base64编码后长度会增加约33%，这里先预估一下
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:length], nil // 截取到指定长度
}
