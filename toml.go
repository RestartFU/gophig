package gophig

import (
	"github.com/pelletier/go-toml"
)

// TOMLMarshaler is a Marshaler that uses the pelletier/go-toml package.
type TOMLMarshaler struct{}

// Marshal ...
func (TOMLMarshaler) Marshal(v interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

// Unmarshal ...
func (TOMLMarshaler) Unmarshal(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}
