package gophig

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
)

func SaveConf[T any](path string, marshaler Marshaler, v T) (T, error) {
	ctx := newRawContext(path, marshaler, os.ModePerm, v)
	return LoadConfContext[T](ctx)
}

func LoadConf[T any](path string, marshaler Marshaler) (T, error) {
	ctx := newRawContext(path, marshaler, os.ModePerm, nil)
	return LoadConfContext[T](ctx)
}

type RawContext struct {
	context.Context
	values map[any]any
}

func newRawContext(name string, marshaler Marshaler, perm os.FileMode, value any) *RawContext {
	return &RawContext{
		Context: context.Background(),
		values: map[any]any{
			"name":      name,
			"marshaler": marshaler,
			"perm":      perm,
			"value":     value,
		},
	}
}

func (c *RawContext) Value(key any) any {
	return c.values[key]
}

// LoadConfContext loads the configuration file into the given type.
func LoadConfContext[T any](ctx context.Context) (T, error) {
	v := new(T)
	name, marshaler, err := extractContextValues(ctx)
	if err != nil {
		return *v, err
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return *v, err
	}

	err = marshaler.Unmarshal(data, v)
	return *v, err
}

// SaveConfContext saves the given type to the configuration file.
func SaveConfContext(ctx context.Context) error {
	name, marshaler, err := extractContextValues(ctx)
	if err != nil {
		return err
	}

	perm, ok := ctx.Value("perm").(os.FileMode)
	if !ok {
		return errors.New("perm not found in context")
	}

	v := ctx.Value("value")
	data, err := marshaler.Marshal(v)
	if err != nil {
		return err
	}

	return os.WriteFile(name, data, perm)
}

func extractContextValues(ctx context.Context) (string, Marshaler, error) {
	var missing []string
	name, ok := ctx.Value("name").(string)
	if !ok {
		missing = append(missing, "name")
	}

	marshaler, ok := ctx.Value("marshaler").(Marshaler)
	if !ok {
		missing = append(missing, "marshaler")
	}

	if len(missing) > 0 {
		return "", marshaler, errors.New(fmt.Sprintf("missing required values in context: %s", strings.Join(missing, ",")))
	}
	return name, marshaler, nil
}
