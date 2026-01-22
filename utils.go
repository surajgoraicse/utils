package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
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
func IsTrue(v bool) bool {
	return v
}
func ParseInt(s string) (int, error) {
	v, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(v), nil

}
