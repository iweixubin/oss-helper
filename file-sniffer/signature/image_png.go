package signature

func init() {
	contentType := "image/png"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"png"},
	}
}
