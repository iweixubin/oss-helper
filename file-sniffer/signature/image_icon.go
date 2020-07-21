package signature

func init() {
	contentType := "image/x-icon"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"icon"},
	}
}
