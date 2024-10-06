package gophig

import (
	"io/fs"
)

// Gophig is a struct that contains the name, extension, and permission of a configuration file.
type Gophig[T any] struct {
	ctx *RawContext

	name      string
	marshaler Marshaler
	perm      fs.FileMode
}

// NewGophig returns a new Gophig struct.
func NewGophig[T any](name string, marshaler Marshaler, perm fs.FileMode) *Gophig[T] {
	g := &Gophig[T]{
		name:      name,
		marshaler: marshaler,
		perm:      perm,

		ctx: newRawContext(name, marshaler, perm, nil),
	}
	return g
}

// WriteConf saves the given interface to the configuration file.
func (g *Gophig[T]) WriteConf(v T) error {
	g.ctx.values["value"] = v
	return WriteConfContext(g.ctx)
}

// ReadConf loads the configuration file into the given interface.
func (g *Gophig[T]) ReadConf() (T, error) {
	return GetConfContext[T](g.ctx)
}
