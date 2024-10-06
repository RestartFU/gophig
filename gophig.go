package gophig

import (
	"context"
	"io/fs"
)

// Gophig is a struct that contains the name, extension, and permission of a configuration file.
type Gophig[T any] struct {
	context.Context

	name      string
	marshaler Marshaler
	perm      fs.FileMode

	value T
}

// NewGophig returns a new Gophig struct.
func NewGophig[T any](name string, marshaler Marshaler, perm fs.FileMode) *Gophig[T] {
	g := &Gophig[T]{
		name:      name,
		marshaler: marshaler,
		perm:      perm,

		Context: context.Background(),
	}
	return g
}

func (g *Gophig[T]) Value(k any) any {
	if k == "name" {
		return g.name
	}
	if k == "marshaler" {
		return g.marshaler
	}
	if k == "perm" {
		return g.perm
	}
	if k == "value" {
		return g.value
	}

	return nil
}

// SetConf saves the given interface to the configuration file.
func (g *Gophig[T]) SetConf(v T) error {
	g.value = v
	return SetConfContext(g)
}

// GetConf loads the configuration file into the given interface.
func (g *Gophig[T]) GetConf() (T, error) {
	return GetConfContext[T](g)
}
