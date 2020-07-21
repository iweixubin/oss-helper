package signature

func init() {
	contentType := "image/bmp"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"bmp"},
	}
}
