package clarity

import "encoding/json"

// Convert interface type to the correct type
func IConvert(data interface{}, v interface{}) error {
	byteData, _ := json.Marshal(data)
	err := json.Unmarshal(byteData, v)
	return err
}
