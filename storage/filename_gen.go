package storage

import (
	"github.com/google/uuid"
	"strings"
)

// genFilename 生产文件名，guid 没有 -，全小写
func genFilename(ext string) string {
	if ext[0] != 46 { // 如果第一个字符不是 .
		ext = "." + ext
	}
	s := uuid.New().String() + ext
	s = strings.Replace(s, "-", "", -1)
	return s
}
