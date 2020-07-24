package signature

func init() {
	contentType := "application/x-rar-compressed"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"rar"},
	}
}
