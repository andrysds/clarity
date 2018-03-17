package clarity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuthorization(t *testing.T) {
	var result error
	var authorization string
	username, password := "clarity", "clarity"

	// invalid authorization length
	authorization = "Basic"
	result = BasicAuthorization(authorization, username, password)
	assert.NotNil(t, result)

	// invalid authorization
	authorization = "Basic Invalid"
	result = BasicAuthorization(authorization, username, password)
	assert.NotNil(t, result)

	// invalid credential
	authorization = "Basic Y2xhcml0eTprbGFyaXRp"
	result = BasicAuthorization(authorization, username, password)
	assert.NotNil(t, result)

	// valid authorization
	authorization = "Basic Y2xhcml0eTpjbGFyaXR5"
	result = BasicAuthorization(authorization, username, password)
	assert.Nil(t, result)
}
