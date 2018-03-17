package clarity

import "testing"

func TestBasicAuthorization(t *testing.T) {
	var authorization string
	username, password := "clarity", "clarity"

	// invalid authorization length
	authorization = "Basic"
	if BasicAuthorization(authorization, username, password) == nil {
		t.Error("should be not nil")
	}

	// invalid authorization
	authorization = "Basic Invalid"
	if BasicAuthorization(authorization, username, password) == nil {
		t.Error("should be not nil")
	}

	// invalid credential
	authorization = "Basic Y2xhcml0eTprbGFyaXRp"
	if BasicAuthorization(authorization, username, password) == nil {
		t.Error("should be not nil")
	}

	// valid authorization
	authorization = "Basic Y2xhcml0eTpjbGFyaXR5"
	if BasicAuthorization(authorization, username, password) != nil {
		t.Error("should be nil")
	}
}
