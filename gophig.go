package gophig

import (
	"io/fs"
)

// Gophig is a struct that contains the name, extension, and permission of a configuration file.
type Gophig struct {
	name      string
	marshaler Marshaler
	perm      fs.FileMode
}

// NewGophig returns a new Gophig struct.
func NewGophig(name string, marshaler Marshaler, perm fs.FileMode) *Gophig {
	return &Gophig{
		name:      name,
		marshaler: marshaler,
		perm:      perm,
	}
}

// SetConf saves the given interface to the configuration file.
func (g *Gophig) SetConf(v interface{}) error {
	return SetConfComplex(g.name, g.marshaler, v, g.perm)
}

// GetConf loads the configuration file into the given interface.
func (g *Gophig) GetConf(v interface{}) error {
	return GetConfComplex(g.name, g.marshaler, v)
}
