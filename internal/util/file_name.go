package util

import (
	"github.com/google/uuid"
	"strings"
	"sync"
	"time"
)

var mutexTimeName sync.Mutex
var lastXSecond int64

// NameByExpr 用过期时间转成一个整数
func NameByExpr(expr time.Time) int64 {
	mutexTimeName.Lock()
	defer mutexTimeName.Unlock()

	year, month, day := expr.Date()
	date := int64(year*1e4+int(month)*1e2+day) * 1e11
	// uint64 最大值是 18446744073709551615 因为年是 2xxx 所以只取 19位

	hour, min, sec := expr.Clock()
	clock := int64(hour*1e4+min*1e2+sec) * 1e5

	for {
		//                 毫秒 微秒 纳秒
		// yyyyMMdd HHmmss.XXX YYY ZZZ
		xSecond := int64(time.Now().Nanosecond() / 1e5)
		// 相当于 5位 随机数

		if xSecond != lastXSecond {
			lastXSecond = xSecond
			break
		}
	}

	return date + clock + lastXSecond
}

//  guid 没有 -，全小写
func GuidName() string {
	s := uuid.New().String()
	s = strings.Replace(s, "-", "", -1)
	return s
}

// NameJoinExt 文件名拼接扩展名
func NameJoinExt(name, ext string) string {
	if ext[0] != 46 { // 如果第一个字符不是 .
		ext = "." + ext
	}

	return name + ext
}
