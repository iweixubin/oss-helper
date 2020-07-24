package signature

func init() {
	contentType := "video/webm"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions:  []string{"webm", "mkv", "mka", "mks", "mk3d"},
	}
}
