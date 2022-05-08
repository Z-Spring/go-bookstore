package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func EncodeMd5(target string) string {
	md5 := md5.New()
	io.WriteString(md5, target)
	return hex.EncodeToString(md5.Sum(nil))
}
