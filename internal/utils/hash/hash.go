//go:build !feature_test

package hash

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func Make(s string) string {
	// ランダムな12文字のsaltを生成
	salt := make([]rune, 12)
	for i := 0; i < 12; i++ {
		v, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		salt[i] = chars[v.Int64()]
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s+string(salt))))
}
