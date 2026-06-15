package gophig

import (
	"errors"
	"fmt"
	"strings"

	"github.com/restartfu/gophig/codecs"
)

// UnsupportedExtensionError is an error that is returned when a file extension is not supported.
type UnsupportedExtensionError struct {
	extension string
}

// Error ...
func (e UnsupportedExtensionError) Error() string {
	return fmt.Sprintf("error: gophig does not support the file extension '%s' at the moment", e.extension)
}

// IsUnsupportedExtensionErr returns true if the error is an UnsupportedExtensionError.
func IsUnsupportedExtensionErr(err error) bool {
	var unsupportedExtensionError UnsupportedExtensionError
	ok := errors.As(err, &unsupportedExtensionError)
	return ok
}

// Marshaler is an interface that can marshal and unmarshal data.
type Marshaler interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

// MarshalerFromExtension is a Marshaler that uses a file extension to determine which Marshaler to use.
func MarshalerFromExtension(ext string) (Marshaler, error) {
	ext = strings.ToLower(ext)
	switch ext {
	case "toml":
		return codecs.TOMLMarshaler{}, nil
	case "json":
		return codecs.JSONMarshaler{}, nil
	case "yaml":
		return codecs.YAMLMarshaler{}, nil
	case "env":
		return codecs.DotenvMarshaler{}, nil
	}
	return nil, UnsupportedExtensionError{ext}
}
