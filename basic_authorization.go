package clarity

import (
	"encoding/base64"
	"errors"
	"strings"
)

func BasicAuthorization(authorization, username, password string) error {
	token := strings.SplitN(authorization, " ", 2)
	if len(token) != 2 || token[0] != "Basic" {
		return errors.New("Invalid Token")
	}

	decoded, err := base64.StdEncoding.DecodeString(token[1])
	if err != nil {
		return err
	}

	credential := strings.SplitN(string(decoded), ":", 2)
	if len(credential) != 2 || credential[0] != username ||
		credential[1] != password {
		return errors.New("Wrong Credential")
	}

	return nil
}
