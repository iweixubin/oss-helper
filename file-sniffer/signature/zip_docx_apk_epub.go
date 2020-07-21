package signature

func init() {
	contentType := "application/zip"

	mapContentType[contentType] = SniffResult{
		ContentType: contentType,
		Extensions: []string{"zip", "aar", "apk", "docx", "epub", "ipa", "jar", "kmz",
			"maff", "odp", "ods", "odt", "pk3", "pk4", "pptx", "usdz", "vsdx", "xlsx", "xpi"},
	}
}
