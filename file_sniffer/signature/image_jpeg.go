package signature

import (
	"oss-helper/enums"
)

func init() {
	contentType := "image/jpeg"

	mapContentType[contentType] = SniffResult{
		FileType:    enums.FileTypeImage,
		UTI:         "public.jpeg",
		Signature:   "jpg", // jpg 与 jpeg 都很常见，取短的~
		ContentType: contentType,
		Extensions:  []string{"jpg", "jpeg", "jpe", "jif", "jfif", "jfi"},
	}
}

// https://en.wikipedia.org/wiki/JPEG
