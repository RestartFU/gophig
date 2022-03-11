package gophig

import (
	"github.com/go-yaml/yaml"
)

type YAMLMarshaler struct{}

func (YAMLMarshaler) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (YAMLMarshaler) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
