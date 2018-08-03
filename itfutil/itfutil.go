package itfutil

import "encoding/json"

// Convert interface type to the correct type
func Convert(data interface{}, v interface{}) error {
	byteData, _ := json.Marshal(data)
	err := json.Unmarshal(byteData, v)
	return err
}
