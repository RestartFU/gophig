package gophig

import "github.com/goccy/go-json"

// JSONMarshaler is a Marshaler that uses the goccy/go-json package.
type JSONMarshaler struct {
	Indent bool
}

// Marshal ...
func (m JSONMarshaler) Marshal(v interface{}) ([]byte, error) {
	if m.Indent {
		return json.MarshalIndent(v, "", "\t")
	}
	return json.Marshal(v)
}

// Unmarshal ...
func (JSONMarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
