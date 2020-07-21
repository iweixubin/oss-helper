package signature

func init() {
	contentType := "image/gif"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"gif"},
	}
}
