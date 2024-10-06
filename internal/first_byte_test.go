package internal_test

import (
	"fmt"
	. "github.com/brownian-motion/rustgo-firstbyte-demo/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FirstByte(t *testing.T) {
	for _, testCase := range []struct {
		input    string
		expected rune
	}{
		{input: "", expected: rune(0)},
		{input: "Some input", expected: rune('S')},
	} {
		t.Run(fmt.Sprintf("\"%s\" -> '%s'", testCase.input, string(testCase.expected)), func(t *testing.T) {
			assert.Equal(t, testCase.expected, FirstByte(testCase.input))
		})
	}
}
