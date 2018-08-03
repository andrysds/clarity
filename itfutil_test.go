package clarity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	var source, result map[string]map[string]int
	source = map[string]map[string]int{
		"parent1": map[string]int{"child": 0},
		"parent2": map[string]int{"child": 1},
	}
	assert.NotEqual(t, source, result)

	// build interface variable
	var itf interface{}
	byteData, _ := json.Marshal(source)
	json.Unmarshal(byteData, &itf)

	// converting
	Convert(itf, &result)
	assert.Equal(t, source, result)
}
