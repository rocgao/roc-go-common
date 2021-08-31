// random
//
// Authors: roc
// Created date: 2021/9/1 00:12

package strext

import "github.com/thanhpk/randstr"

// NewSalt 生成20位的随机盐
func NewSalt() string {
	return randstr.String(20)
}

// NewAesKey 生成24位的随机AesKey
func NewAesKey() string {
	return randstr.String(24)
}

// NewTraceID 生成12位的随机traceID
func NewTraceID() string {
	return randstr.String(12)
}
