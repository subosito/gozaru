package norma

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	NormalizationTable = map[string][]string{
		"a":   []string{"a", " a", "a ", " a ", "a    \n"},
		"x x": []string{"x x", "x  x", "x   x", " x  |  x ", "x\tx", "x\r\nx"},
	}

	SanitizationTable = []string{"<", ">", "|", "/", "\\", "*", "?", ":"}
)

func TestNorma_normalization(t *testing.T) {
	for k, v := range NormalizationTable {
		for i := range v {
			assert.Equal(t, k, Sanitize(v[i]))
		}
	}
}

func TestNorma_truncation(t *testing.T) {
	name := strings.Repeat("A", 400)
	assert.Equal(t, 255, len(Sanitize(name)))
	assert.Equal(t, 245, len(SanitizePad(name, 10)))
}

func TestNorma_sanitization(t *testing.T) {
	assert.Equal(t, "abcdef", Sanitize(`abcdef`))

	for i := range SanitizationTable {
		k := SanitizationTable[i]

		assert.Equal(t, "file", Sanitize(k))
		assert.Equal(t, "a", Sanitize(fmt.Sprintf("a%s", k)))
		assert.Equal(t, "a", Sanitize(fmt.Sprintf("%sa", k)))
		assert.Equal(t, "aa", Sanitize(fmt.Sprintf("a%sa", k)))
	}

	assert.Equal(t, `笊, ざる.pdf`, Sanitize(`笊, ざる.pdf`))
	assert.Equal(t, `whatēverwëirduserînput`, Sanitize(`  what\\ēver//wëird:user:înput:`))
}

func TestNorma_windowsReservedNamed(t *testing.T) {
	assert.Equal(t, "file", Sanitize(`CON`))
	assert.Equal(t, "file", Sanitize(`lpt1`))
	assert.Equal(t, "file", Sanitize(`com4`))
	assert.Equal(t, "file", Sanitize(` aux`))
	assert.Equal(t, "file", Sanitize(" LpT\x122"))
	assert.Equal(t, "COM10", Sanitize(`COM10`))
}

func TestNorma_blanks(t *testing.T) {
	assert.Equal(t, "file", Sanitize(`<`))
}

func TestNorma_dots(t *testing.T) {
	assert.Equal(t, "file.pdf", Sanitize(`.pdf`))
	assert.Equal(t, "file.pdf", Sanitize(`<.pdf`))
	assert.Equal(t, "file..pdf", Sanitize(`..pdf`))
}

func TestNorma_fallback(t *testing.T) {
	assert.Equal(t, "blub", SanitizeFallback(`<`, "blub"))
	assert.Equal(t, "blub", SanitizeFallback(`lpt1`, "blub"))
	assert.Equal(t, "blub.pdf", SanitizeFallback(`<.pdf`, "blub"))
}
