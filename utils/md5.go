package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func MD5(s string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, s)
	return hex.EncodeToString(hash.Sum(nil))
}
