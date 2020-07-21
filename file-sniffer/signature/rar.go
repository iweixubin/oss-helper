package signature

func init() {
	contentType := "application/x-gzip"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"gz", "tar.gz"},
	}
}
