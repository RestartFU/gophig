package gophig

import (
	"io/fs"
	"os"
)

// GetConfComplex loads the configuration file into the given interface.
func GetConfComplex(name string, marshaler Marshaler, v any) error {
	data, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	return marshaler.Unmarshal(data, v)
}

// SetConfComplex saves the given interface to the configuration file.
func SetConfComplex(name string, marshaler Marshaler, v any, perm fs.FileMode) error {
	data, err := marshaler.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(name, data, perm)
}
