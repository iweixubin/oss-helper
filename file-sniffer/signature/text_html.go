package signature

func init() {
	contentType := "text/html; charset=utf-8"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"html"},
	}
}
