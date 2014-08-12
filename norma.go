package norma

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	CharacterFilter   = `[\x00-\x1F\/\\:\*\?\"<>\|]`
	UnicodeWhitespace = `[[:space:]]+`
)

var (
	FallbackFilename     = "file"
	WindowsReservedNames = [...]string{
		"CON", "PRN", "AUX", "NUL",
		"COM1", "COM2", "COM3", "COM4", "COM5",
		"COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5",
		"LPT6", "LPT7", "LPT8", "LPT9",
	}
)

func Sanitize(s string) string {
	return truncate(s, 0)
}

func SanitizePad(s string, n int) string {
	return truncate(s, n)
}

func normalize(s string) string {
	rx := regexp.MustCompile(UnicodeWhitespace)
	return strings.TrimSpace(rx.ReplaceAllString(s, " "))
}

func sanitize(s string) string {
	rx := regexp.MustCompile(CharacterFilter)
	sc := normalize(s)
	sc = strings.TrimSpace(rx.ReplaceAllString(sc, ""))

	return filter(sc)
}

func truncate(s string, n int) string {
	sc := sanitize(s)
	nc := len(sc)

	if nc > 255 {
		nc = 255
	}

	if n != 0 {
		nc -= n
	}

	return sc[0:nc]
}

func filter(s string) string {
	s = filterWindowsReservedNames(s)
	s = filterBlank(s)
	s = filterDot(s)
	return s
}

func filterWindowsReservedNames(s string) string {
	us := strings.ToUpper(s)

	for i := range WindowsReservedNames {
		v := WindowsReservedNames[i]

		if v == us {
			return FallbackFilename
		}
	}

	return s
}

func filterBlank(s string) string {
	if s == "" {
		return FallbackFilename
	}

	return s
}

func filterDot(s string) string {
	if strings.HasPrefix(s, ".") {
		return fmt.Sprintf("%s%s", FallbackFilename, s)
	}

	return s
}
