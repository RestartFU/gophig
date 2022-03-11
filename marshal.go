package gophig

import (
	"fmt"
	"strings"
)

type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

func extMarshaler(ext string) (Marshaler, error) {
	ext = strings.ToLower(ext)
	switch ext {
	case "toml":
		return TOMLMarshaler{}, nil
	case "json":
		return JSONMarshaler{}, nil
	case "yaml":
		return YAMLMarshaler{}, nil
	}
	return nil, fmt.Errorf("error: gophig does not support the file extension '%s' at the moment", ext)
}
