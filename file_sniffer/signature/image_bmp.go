package signature

import (
	"oss-helper/enums"
)

func init() {
	contentType := "image/bmp"
	UTI := "com.microsoft.bmp"
	Signature := "bmp"
	Extensions := []string{"bmp", "dib"}

	mapContentType[contentType] = SniffResult{
		FileType:    enums.FileTypeImage,
		UTI:         UTI,
		Signature:   Signature,
		ContentType: contentType,
		Extensions:  Extensions,
	}

	// 有另外一种
	contentTypeX := "image/x-bmp"

	mapContentType[contentTypeX] = SniffResult{
		FileType:    enums.FileTypeImage,
		UTI:         UTI,
		Signature:   Signature,
		ContentType: contentTypeX,
		Extensions:  Extensions,
	}
}

// https://en.wikipedia.org/wiki/BMP_file_format
