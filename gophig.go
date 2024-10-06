package gophig

import (
	"io/fs"
)

// Gophig is a struct that contains the name, extension, and permission of a configuration file.
type Gophig[T any] struct {
	name      string
	marshaler Marshaler
	perm      fs.FileMode
}

// NewGophig returns a new Gophig struct.
func NewGophig[T any](name string, marshaler Marshaler, perm fs.FileMode) *Gophig[T] {
	return &Gophig[T]{
		name:      name,
		marshaler: marshaler,
		perm:      perm,
	}
}

// SetConf saves the given interface to the configuration file.
func (g *Gophig[T]) SetConf(v T) error {
	return SetConfComplex(g.name, g.marshaler, v, g.perm)
}

// GetConf loads the configuration file into the given interface.
func (g *Gophig[T]) GetConf() (T, error) {
	v := new(T)
	v2 := *v
	err := GetConfComplex(g.name, g.marshaler, &v2)
	return v2, err
}
