# Gophig

Gophig is a simple configuration manager for Go projects. It supports marshaling and unmarshaling config files (e.g., TOML) into typed Go structs.

# Installation

```sh
go get github.com/restartfu/gophig
```

# Project structure

```text
.
|-- config.go              # load/save functions
|-- gophig.go              # typed config manager
|-- marshal.go             # Marshaler interface and extension lookup
|-- codecs/                # JSON, TOML, YAML, dotenv marshalers
|-- tests/                 # external package tests and fixtures
`-- examples/basic/        # runnable example
```

# Usage

Define your configuration struct:

```go
import (
	"github.com/restartfu/gophig"
	"github.com/restartfu/gophig/codecs"
)

type Foo struct {
	Foo string `toml:"foo"`
	Bar string `toml:"bar"`
}
```

# Create a new *Gophig instance:
```go
g := gophig.NewGophig[Foo]("./config.toml", codecs.TOMLMarshaler{}, os.ModePerm)
```
# Writing a Config
```go
myFoo := Foo{Foo: "foo", Bar: "bar"}

if err := g.SaveConf(myFoo); err != nil {
	log.Fatalln(err)
}
```

This will generate a config.toml file with:
```
foo = "foo"
bar = "bar"
```
# Reading a Config

```go
myFoo, err := g.LoadConf()
if err != nil {
  log.Fatalln(err)
}

log.Println(myFoo)
```

Output:

```
{Foo: "foo", Bar: "bar"}
```

# Environment Variables

Config values can reference environment variables with `${VAR}`. Expansion happens before unmarshaling. Missing variables are left unchanged unless a default is provided with `${VAR:-default}`.

```toml
host = "${APP_HOST}"
port = "${APP_PORT:-8080}"
```
