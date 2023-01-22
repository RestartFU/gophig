package gophig

import "encoding/json"

// JSONMarshaler is a Marshaler that uses the standard library's json package.
type JSONMarshaler struct{}

// Marshal ...
func (JSONMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal ...
func (JSONMarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
