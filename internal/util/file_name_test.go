package util

import (
	"testing"
	"time"
)

func Test_TimeName(t *testing.T) {
	TimeName()
	tt := time.Now()
	println(tt.Format("20060102150405.000000"))
	println(tt.Nanosecond())
}
