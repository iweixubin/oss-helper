package signature

import (
	"oss-helper/enums"
)

func init() {
	contentType := "image/webp"

	mapContentType[contentType] = SniffResult{
		FileType:    enums.FileTypeImage,
		UTI:         "com.google.webp", // wiki 上没有，但参考 bmp 的规则，应该是这样~
		Signature:   "webp",
		ContentType: contentType,
		Extensions:  []string{"webp"},
	}
}

// https://en.wikipedia.org/wiki/WebP
