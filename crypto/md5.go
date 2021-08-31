// md5
//
// Authors: roc
// Created date: 2021/9/1 00:14

package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// CalcMD5 计算text的MD5值
func CalcMD5(text string) string {
	if text == "" {
		return ""
	}

	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
