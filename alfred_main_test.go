package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	assert.Equal(t, 12, lengthOf("hello, world", lenByRunes))
	assert.Equal(t, 3, lengthOf("123", lenByRunes))
	assert.Equal(t, 4, lengthOf("你好世界", lenByRunes))

	assert.Equal(t, 12, lengthOf("hello, world", lenBySize))
	assert.Equal(t, 3, lengthOf("123", lenBySize))
	assert.Equal(t, 12, lengthOf("hello 你好", lenBySize))
	assert.Equal(t, 12, lengthOf("你好世界", lenBySize))
}
