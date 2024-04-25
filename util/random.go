package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func RandInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func RandString(length int) string {
	var sb strings.Builder
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(byte(RandInt('a', 'z')))
	}
	return sb.String()
}
func RandomOwner() string {
	return RandString(6)
}
func RandomMoney() int64 {
	return RandInt(0, 1000)
}
func RandomCurrency() string {
	curruncies := []string{"EUR", "USD", "GBP", "JPY"}
	n := len(curruncies)
	return curruncies[rand.Intn(n)]
}
