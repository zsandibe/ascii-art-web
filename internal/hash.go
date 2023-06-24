package internal

import (
	"crypto/md5"
	"fmt"
)

func HashMD5(content string) string {
	h := md5.Sum([]byte(content))
	return fmt.Sprintf("%x", h)
}

func CheckHash(hash string) bool {
	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy := "86d9947457f6a41a18cb98427e314ff8"
	if hash == hashStandard {
		return true
	}
	if hash == hashShadow {
		return true
	}
	if hash == hashThinkertoy {
		return true
	}
	return false
}
