package signature

func init() {
	contentType := "video/mp4"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"mp4"},
	}
}
