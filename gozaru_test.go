package gozaru_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subosito/gozaru"
)

var (
	GozarulizationTable = map[string][]string{
		"a":   []string{"a", " a", "a ", " a ", "a    \n"},
		"x x": []string{"x x", "x  x", "x   x", " x  |  x ", "x\tx", "x\r\nx"},
	}

	SanitizationTable = []string{"<", ">", "|", "/", "\\", "*", "?", ":"}
)

func TestGozaru_normalization(t *testing.T) {
	for k, v := range GozarulizationTable {
		for i := range v {
			assert.Equal(t, k, gozaru.Sanitize(v[i]))
		}
	}
}

func TestGozaru_truncation(t *testing.T) {
	name := strings.Repeat("A", 400)
	assert.Equal(t, 255, len(gozaru.Sanitize(name)))
	assert.Equal(t, 245, len(gozaru.SanitizePad(name, 10)))
}

func TestGozaru_sanitization(t *testing.T) {
	assert.Equal(t, "abcdef", gozaru.Sanitize(`abcdef`))

	for i := range SanitizationTable {
		k := SanitizationTable[i]

		assert.Equal(t, "file", gozaru.Sanitize(k))
		assert.Equal(t, "a", gozaru.Sanitize(fmt.Sprintf("a%s", k)))
		assert.Equal(t, "a", gozaru.Sanitize(fmt.Sprintf("%sa", k)))
		assert.Equal(t, "aa", gozaru.Sanitize(fmt.Sprintf("a%sa", k)))
	}

	assert.Equal(t, `笊, ざる.pdf`, gozaru.Sanitize(`笊, ざる.pdf`))
	assert.Equal(t, `whatēverwëirduserînput`, gozaru.Sanitize(`  what\\ēver//wëird:user:înput:`))
}

func TestGozaru_windowsReservedNamed(t *testing.T) {
	assert.Equal(t, "file", gozaru.Sanitize(`CON`))
	assert.Equal(t, "file", gozaru.Sanitize(`lpt1`))
	assert.Equal(t, "file", gozaru.Sanitize(`com4`))
	assert.Equal(t, "file", gozaru.Sanitize(` aux`))
	assert.Equal(t, "file", gozaru.Sanitize(" LpT\x122"))
	assert.Equal(t, "COM10", gozaru.Sanitize(`COM10`))
}

func TestGozaru_blanks(t *testing.T) {
	assert.Equal(t, "file", gozaru.Sanitize(`<`))
}

func TestGozaru_dots(t *testing.T) {
	assert.Equal(t, "file.pdf", gozaru.Sanitize(`.pdf`))
	assert.Equal(t, "file.pdf", gozaru.Sanitize(`<.pdf`))
	assert.Equal(t, "file..pdf", gozaru.Sanitize(`..pdf`))
}

func TestGozaru_fallback(t *testing.T) {
	assert.Equal(t, "blub", gozaru.SanitizeFallback(`<`, "blub"))
	assert.Equal(t, "blub", gozaru.SanitizeFallback(`lpt1`, "blub"))
	assert.Equal(t, "blub.pdf", gozaru.SanitizeFallback(`<.pdf`, "blub"))

	name := strings.Repeat(" ", 400)
	assert.Equal(t, "blub", gozaru.SanitizeFallback(name, "blub"))
	assert.Equal(t, "blub", gozaru.SanitizePadFallback(name, 10, "blub"))
}
