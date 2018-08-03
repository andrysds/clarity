package clarity

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

// BasicAuthorizer is where you store credential for basic authorization
type BasicAuthorizer struct {
	Username string
	Password string
}

// NewBasicAuthorizer prepare an instance of BasicAuthorizer
func NewBasicAuthorizer(username, password string) *BasicAuthorizer {
	return &BasicAuthorizer{
		Username: username,
		Password: password,
	}
}

// Authorize do basic authorization
func (b *BasicAuthorizer) Authorize(header http.Header) error {
	authorization := header.Get("Authorization")
	token := strings.SplitN(authorization, " ", 2)
	if len(token) != 2 || token[0] != "Basic" {
		return errors.New("Invalid Token")
	}

	decoded, err := base64.StdEncoding.DecodeString(token[1])
	if err != nil {
		return err
	}

	credential := strings.SplitN(string(decoded), ":", 2)
	if len(credential) != 2 || credential[0] != b.Username ||
		credential[1] != b.Password {
		return errors.New("Wrong Credential")
	}

	return nil
}
