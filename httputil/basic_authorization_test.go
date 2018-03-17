package httputil

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuthorization(t *testing.T) {
	var result error
	header := http.Header{}
	username, password := "clarity", "clarity"

	// invalid authorization length
	header.Set("Authorization", "Basic")
	result = BasicAuthorization(header, username, password)
	assert.NotNil(t, result)

	// invalid authorization
	header.Set("Authorization", "Basic Invalid")
	result = BasicAuthorization(header, username, password)
	assert.NotNil(t, result)

	// invalid credential
	header.Set("Authorization", "Basic Y2xhcml0eTprbGFyaXRp")
	result = BasicAuthorization(header, username, password)
	assert.NotNil(t, result)

	// valid authorization
	header.Set("Authorization", "Basic Y2xhcml0eTpjbGFyaXR5")
	result = BasicAuthorization(header, username, password)
	assert.Nil(t, result)
}

func ExampleBasicAuthorization() {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	var header http.Header
	header.Set("Authorization", "Basic Y2xhcml0eTpjbGFyaXR5")
	err := BasicAuthorization(header, username, password)
	fmt.Println(err)
}
