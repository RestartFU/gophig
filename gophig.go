package gophig

import (
	"io/fs"
)

// Gophig is a struct that contains the name, extension, and permission of a configuration file.
type Gophig struct {
	Name, Extension string
	Perm            fs.FileMode
}

// NewGophig returns a new Gophig struct.
func NewGophig(name, extension string, perm fs.FileMode) *Gophig {
	return &Gophig{
		Name:      name,
		Extension: extension,
		Perm:      perm,
	}
}

// SetConf saves the given interface to the configuration file.
func (gophig *Gophig) SetConf(v interface{}) error {
	marshaler, err := extMarshaler(gophig.Extension)
	if err != nil {
		return err
	}
	return SetConfComplex(gophig.Name+"."+gophig.Extension, marshaler, v, gophig.Perm)
}

// GetConf loads the configuration file into the given interface.
func (gophig *Gophig) GetConf(v interface{}) error {
	marshaler, err := extMarshaler(gophig.Extension)
	if err != nil {
		return err
	}
	return GetConfComplex(gophig.Name+"."+gophig.Extension, marshaler, v)
}
