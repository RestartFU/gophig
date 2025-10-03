package gophig

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

type DotenvMarshaler struct{}

func (DotenvMarshaler) Marshal(v any) ([]byte, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return nil, fmt.Errorf("DotenvMarshaler: Marshal expects a valid value")
	}

	var sb strings.Builder

	switch rv.Kind() {
	case reflect.Map:
		m, ok := v.(map[string]string)
		if !ok {
			return nil, fmt.Errorf("DotenvMarshaler: Marshal expects map[string]string")
		}
		for key, val := range m {
			val = strings.ReplaceAll(val, "\n", `\n`)
			sb.WriteString(fmt.Sprintf("%s=%s\n", key, val))
		}

	case reflect.Struct:
		rt := rv.Type()
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			fv := rv.Field(i)
			if !fv.CanInterface() {
				continue
			}

			key := field.Tag.Get("env")
			if key == "" {
				continue // skip fields without env tag
			}

			val := fmt.Sprintf("%v", fv.Interface())
			val = strings.ReplaceAll(val, "\n", `\n`)
			sb.WriteString(fmt.Sprintf("%s=%s\n", key, val))
		}

	default:
		return nil, fmt.Errorf("DotenvMarshaler: unsupported type %s", rv.Kind())
	}

	return []byte(sb.String()), nil
}

func (DotenvMarshaler) Unmarshal(data []byte, v any) error {
	envMap, err := godotenv.Unmarshal(string(data))
	if err != nil {
		return fmt.Errorf("DotenvMarshaler: failed to unmarshal dotenv: %w", err)
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return fmt.Errorf("DotenvMarshaler: Unmarshal expects a non-nil pointer")
	}

	rvElem := rv.Elem()

	switch rvElem.Kind() {
	case reflect.Map:
		if rvElem.Type().Key().Kind() != reflect.String || rvElem.Type().Elem().Kind() != reflect.String {
			return fmt.Errorf("DotenvMarshaler: only map[string]string supported")
		}
		rvElem.Set(reflect.ValueOf(envMap))

	case reflect.Struct:
		for i := 0; i < rvElem.NumField(); i++ {
			field := rvElem.Type().Field(i)
			tag := field.Tag.Get("env")
			if tag == "" {
				tag = strings.ToUpper(field.Name)
			}
			if val, ok := envMap[tag]; ok {
				f := rvElem.Field(i)
				if f.CanSet() {
					switch f.Kind() {
					case reflect.String:
						f.SetString(val)
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						var intVal int64
						_, err := fmt.Sscan(val, &intVal)
						if err != nil {
							return fmt.Errorf("DotenvMarshaler: failed to parse int for field %s: %w", field.Name, err)
						}
						f.SetInt(intVal)
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						var uintVal uint64
						_, err := fmt.Sscan(val, &uintVal)
						if err != nil {
							return fmt.Errorf("DotenvMarshaler: failed to parse uint for field %s: %w", field.Name, err)
						}
						f.SetUint(uintVal)
					case reflect.Bool:
						var boolVal bool
						_, err := fmt.Sscan(val, &boolVal)
						if err != nil {
							return fmt.Errorf("DotenvMarshaler: failed to parse bool for field %s: %w", field.Name, err)
						}
						f.SetBool(boolVal)
					case reflect.Float32, reflect.Float64:
						var floatVal float64
						_, err := fmt.Sscan(val, &floatVal)
						if err != nil {
							return fmt.Errorf("DotenvMarshaler: failed to parse float for field %s: %w", field.Name, err)
						}
						f.SetFloat(floatVal)
					default:
						return fmt.Errorf("DotenvMarshaler: unsupported field type %s for field %s", f.Kind(), field.Name)
					}
				}
			}
		}

	default:
		return fmt.Errorf("DotenvMarshaler: unsupported type %s", rvElem.Kind())
	}

	return nil
}
