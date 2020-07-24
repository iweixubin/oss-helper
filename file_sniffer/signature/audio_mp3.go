package signature

func init() {
	contentType := "audio/mpeg"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"mp3"},
	}
}
