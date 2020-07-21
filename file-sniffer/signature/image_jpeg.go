package signature

func init() {
	contentType := "image/jpeg"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"jpg", "jpeg"},
	}
}
