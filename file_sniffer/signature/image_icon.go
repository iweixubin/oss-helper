package signature

import (
	"oss-helper/enums"
)

func init() {
	contentType := "image/x-icon"

	mapContentType[contentType] = SniffResult{
		FileType:    enums.FileTypeImage,
		UTI:         "com.microsoft.ico",
		Signature:   "ico",
		ContentType: contentType,
		Extensions:  []string{"ico", "icon"},
	}
}

// https://en.wikipedia.org/wiki/ICO_(file_format)
