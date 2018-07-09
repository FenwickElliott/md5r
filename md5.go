package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// String generates and returns a md5 hash of the given string
func String(in string) string {
	h := md5.New()
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))
}
