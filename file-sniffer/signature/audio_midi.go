package signature

func init() {
	contentType := "audio/midi"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"midi,mid"},
	}
}
