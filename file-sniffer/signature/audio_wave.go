package signature

func init() {
	contentType := "audio/wave"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"wav"},
	}
}
