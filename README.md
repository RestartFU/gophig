# Gophig

Gophig is a simple configuration manager for Go projects. It supports marshaling and unmarshaling config files (e.g., TOML) into typed Go structs.

# Installation

go get git.restartfu.com/restart/gophig

# Usage

Define your configuration struct:

```go
type Foo struct {
	Foo string toml:"foo"
	Bar string toml:"bar"
}
```

# Create a new *Gophig instance:
```go
g := gophig.NewGophig[Foo]("./config.toml", gophig.TOMLMarshaler, os.ModePerm)
```
# Writing a Config
```go
myFoo := Foo{Foo: "foo", Bar: "bar"}

if err := g.WriteConf(myFoo); err != nil {
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
myFoo, err := g.ReadConf()
if err != nil {
  log.Fatalln(err)
}

log.Println(myFoo)
```

Output:

```
{Foo: "foo", Bar: "bar"}
```
