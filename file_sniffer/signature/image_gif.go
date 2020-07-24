package signature

import (
	"oss-helper/enums"
)

func init() {
	contentType := "image/gif"

	mapContentType[contentType] = SniffResult{
		FileType:    enums.FileTypeImage,
		UTI:         "com.compuserve.gif",
		Signature:   "gif",
		ContentType: contentType,
		Extensions:  []string{"gif"},
	}
}

// https://en.wikipedia.org/wiki/GIF
