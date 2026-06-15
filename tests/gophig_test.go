package gophig_test

import (
	"os"
	"testing"

	"github.com/restartfu/gophig"
	"github.com/stretchr/testify/require"
)

type MockMarshaler struct{}

func (m MockMarshaler) Marshal(any) ([]byte, error) {
	return []byte{}, nil
}
func (MockMarshaler) Unmarshal([]byte, any) error {
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

type EnvSample struct {
	Host string `env:"HOST"`
}

func TestGophig_GetConf(t *testing.T) {
	for _, ext := range []string{
		"json",
		"toml",
		"yaml",
		"env",
	} {
		t.Run("sample unmarshals successfully into "+ext+" sample struct", func(t *testing.T) {
			marshaler, err := gophig.MarshalerFromExtension(ext)
			require.NoError(t, err)

			g := gophig.NewGophig[Sample]("assets/sample."+ext, marshaler, os.ModePerm)
			require.NotNil(t, g)

			sample, err := g.LoadConf()
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
		"env",
	} {
		t.Run("sample marshals successfully into "+ext+" sample data", func(t *testing.T) {
			marshaler, err := gophig.MarshalerFromExtension(ext)
			require.NoError(t, err)

			g := gophig.NewGophig[Sample](t.TempDir()+"/sample."+ext, marshaler, os.ModePerm)
			require.NotNil(t, g)

			sample := Sample{
				Name:    "jane",
				Surname: "doe",
				Age:     20,
			}

			err = g.SaveConf(sample)
			require.NoError(t, err)
		})
	}
}

func TestGophig_LoadConfExpandsEnvironmentVariables(t *testing.T) {
	t.Setenv("GOPHIG_TEST_HOST", "localhost")

	files := map[string]string{
		"json": `{"Host":"${GOPHIG_TEST_HOST}"}`,
		"toml": `Host = '${GOPHIG_TEST_HOST}'`,
		"yaml": `host: ${GOPHIG_TEST_HOST}`,
		"env":  `HOST=${GOPHIG_TEST_HOST}`,
	}

	for ext, contents := range files {
		t.Run("expands "+ext+" config values", func(t *testing.T) {
			marshaler, err := gophig.MarshalerFromExtension(ext)
			require.NoError(t, err)

			path := t.TempDir() + "/config." + ext
			err = os.WriteFile(path, []byte(contents), 0o644)
			require.NoError(t, err)

			g := gophig.NewGophig[EnvSample](path, marshaler, os.ModePerm)
			sample, err := g.LoadConf()
			require.NoError(t, err)
			require.Equal(t, "localhost", sample.Host)
		})
	}
}

func TestGophig_LoadConfPreservesMissingEnvironmentVariables(t *testing.T) {
	const envName = "GOPHIG_MISSING_HOST_TEST_SHOULD_NOT_EXIST"
	oldValue, wasSet := os.LookupEnv(envName)
	require.NoError(t, os.Unsetenv(envName))
	t.Cleanup(func() {
		if wasSet {
			require.NoError(t, os.Setenv(envName, oldValue))
		}
	})

	path := t.TempDir() + "/config.json"
	err := os.WriteFile(path, []byte(`{"Host":"${GOPHIG_MISSING_HOST_TEST_SHOULD_NOT_EXIST}"}`), 0o644)
	require.NoError(t, err)

	marshaler, err := gophig.MarshalerFromExtension("json")
	require.NoError(t, err)

	g := gophig.NewGophig[EnvSample](path, marshaler, os.ModePerm)
	sample, err := g.LoadConf()
	require.NoError(t, err)
	require.Equal(t, "${"+envName+"}", sample.Host)
}

func TestGophig_LoadConfUsesDefaultForMissingEnvironmentVariables(t *testing.T) {
	const envName = "GOPHIG_DEFAULT_HOST_TEST_SHOULD_NOT_EXIST"
	oldValue, wasSet := os.LookupEnv(envName)
	require.NoError(t, os.Unsetenv(envName))
	t.Cleanup(func() {
		if wasSet {
			require.NoError(t, os.Setenv(envName, oldValue))
		}
	})

	path := t.TempDir() + "/config.json"
	err := os.WriteFile(path, []byte(`{"Host":"${GOPHIG_DEFAULT_HOST_TEST_SHOULD_NOT_EXIST:-localhost}"}`), 0o644)
	require.NoError(t, err)

	marshaler, err := gophig.MarshalerFromExtension("json")
	require.NoError(t, err)

	g := gophig.NewGophig[EnvSample](path, marshaler, os.ModePerm)
	sample, err := g.LoadConf()
	require.NoError(t, err)
	require.Equal(t, "localhost", sample.Host)
}

func TestGophig_LoadConfEnvironmentVariableOverridesDefault(t *testing.T) {
	t.Setenv("GOPHIG_DEFAULT_HOST", "production")

	path := t.TempDir() + "/config.json"
	err := os.WriteFile(path, []byte(`{"Host":"${GOPHIG_DEFAULT_HOST:-localhost}"}`), 0o644)
	require.NoError(t, err)

	marshaler, err := gophig.MarshalerFromExtension("json")
	require.NoError(t, err)

	g := gophig.NewGophig[EnvSample](path, marshaler, os.ModePerm)
	sample, err := g.LoadConf()
	require.NoError(t, err)
	require.Equal(t, "production", sample.Host)
}
