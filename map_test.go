package clarity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapJoin(t *testing.T) {
	a := map[string]string{
		"first name": "andrys",
		"last name":  "silalahi",
	}
	b := map[string]string{
		"github":    "andrysds",
		"instagram": "andrysds_",
	}
	c := map[string]string{
		"first name": "andrys",
		"last name":  "silalahi",
		"github":     "andrysds",
		"instagram":  "andrysds_",
	}
	d := MapJoin(a, b)
	assert.Equal(t, c, d)
}
