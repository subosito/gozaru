package norma

func Sanitize(s string) string {
	return normalize(s)
}

func SanitizePad(s string, n int) string {
	return normalize(s)
}

func normalize(s string) string {
	return s
}
