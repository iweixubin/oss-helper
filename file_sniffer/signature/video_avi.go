package signature

func init() {
	contentType := "video/avi"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"avi"},
	}
}
