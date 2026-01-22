package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func IsEven(num int) bool {
	if num&1 == 0 {
		return true
	}
	return false
}
func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	str := hex.EncodeToString(bs)
	return str
}
