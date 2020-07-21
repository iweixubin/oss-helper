package signature

func init() {
	contentType := "image/webp"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"webp"},
	}
}
