package webart

func MapFont(fontname string) string {
	var fonts = map[string]string{
		"standard":   "../standard.txt",
		"shadow":     "../shadow.txt",
		"thinkertoy": "../thinkertoy.txt",
	}
	return fonts[fontname]
}
