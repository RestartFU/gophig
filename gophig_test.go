package gophig_test

import (
	"github.com/restartfu/gophig"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type MockMarshaler struct{}

func (m MockMarshaler) Marshal(interface{}) ([]byte, error) {
	return []byte{}, nil
}
func (MockMarshaler) Unmarshal([]byte, interface{}) error {
	return nil
}

func TestNewGophig(t *testing.T) {
	t.Run("gophig creation is successful", func(t *testing.T) {
		g := gophig.NewGophig[any]("", MockMarshaler{}, os.ModePerm)
		require.NotNil(t, g)
	})
}

type Sample struct {
	Name    string `json,toml,yaml:"name"`
	Surname string `json,toml,yaml:"surname"`
	Age     int    `json,toml,yaml:"age"`
}

func TestGophig_GetConf(t *testing.T) {
	for _, ext := range []string{
		"json",
		"toml",
		"yaml",
	} {
		t.Run("sample unmarshals successfully into "+ext+" sample struct", func(t *testing.T) {
			marshaler, err := gophig.MarshalerFromExtension(ext)
			require.NoError(t, err)

			g := gophig.NewGophig[Sample]("tests/assets/sample."+ext, marshaler, os.ModePerm)
			require.NotNil(t, g)

			sample, err := g.GetConf()
			require.NoError(t, err)

			require.Equal(t,
				Sample{
					Name:    "jane",
					Surname: "doe",
					Age:     20,
				},
				sample,
			)
		})
	}
}

func TestGophig_SetConf(t *testing.T) {
	for _, ext := range []string{
		"json",
		"toml",
		"yaml",
	} {
		t.Run("sample marshals successfully into "+ext+" sample data", func(t *testing.T) {
			marshaler, err := gophig.MarshalerFromExtension(ext)
			require.NoError(t, err)

			err = os.MkdirAll("tests/tmp", os.ModePerm)
			require.NoError(t, err)
			defer func() {
				require.NoError(t, os.RemoveAll("tests/tmp"))
			}()

			g := gophig.NewGophig[Sample]("tests/tmp/sample."+ext, marshaler, os.ModePerm)
			require.NotNil(t, g)

			sample := Sample{
				Name:    "jane",
				Surname: "doe",
				Age:     20,
			}

			err = g.SetConf(sample)
			require.NoError(t, err)
		})
	}
}
