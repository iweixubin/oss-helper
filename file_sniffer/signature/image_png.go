package signature

import (
	"oss-helper/enums"
)

func init() {
	contentType := "image/png"

	mapContentType[contentType] = SniffResult{
		FileType:    enums.FileTypeImage,
		UTI:         "public.png",
		Signature:   "png",
		ContentType: contentType,
		Extensions:  []string{"png"},
	}
}

// https://en.wikipedia.org/wiki/Portable_Network_Graphics
