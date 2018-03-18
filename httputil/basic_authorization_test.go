package httputil

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBasicAuthorizer(t *testing.T) {
	result := NewBasicAuthorizer("clarity", "clarity")
	assert.NotNil(t, result)
}

func TestAuthorize(t *testing.T) {
	var result error
	authorizer := NewBasicAuthorizer("clarity", "clarity")
	header := http.Header{}

	// invalid authorization length
	header.Set("Authorization", "Basic")
	result = authorizer.Authorize(header)
	assert.NotNil(t, result)

	// invalid authorization
	header.Set("Authorization", "Basic Invalid")
	result = authorizer.Authorize(header)
	assert.NotNil(t, result)

	// invalid credential
	header.Set("Authorization", "Basic Y2xhcml0eTprbGFyaXRp")
	result = authorizer.Authorize(header)
	assert.NotNil(t, result)

	// valid authorization
	header.Set("Authorization", "Basic Y2xhcml0eTpjbGFyaXR5")
	result = authorizer.Authorize(header)
	assert.Nil(t, result)
}
