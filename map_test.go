package clarity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	// error
	assert.Equal(t, 0, Atoi("clarity"))

	// normal
	assert.Equal(t, 1, Atoi("1"))
}
