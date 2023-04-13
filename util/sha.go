package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// GetSha224 密码ssha224加密函数
func GetSha224(str string) string {
	srcByte := []byte(str)
	sha224new := sha256.New224()
	sha224new.Write(srcByte)
	hashed := sha224new.Sum(nil)
	//log.Printf("原密码为:%s 加密后为:%s", str, hex.EncodeToString(hashed))
	return hex.EncodeToString(hashed)
}
