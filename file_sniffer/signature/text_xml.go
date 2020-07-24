package signature

func init() {
	contentType := "text/xml; charset=utf-8"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"xml"},
	}
}
