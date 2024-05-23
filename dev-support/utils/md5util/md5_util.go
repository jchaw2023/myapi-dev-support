package md5util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5str(c string) string {
	_sumBytes := md5.Sum([]byte(c)) //内容key 区分内容是否相同
	return hex.EncodeToString(_sumBytes[0:])
}
