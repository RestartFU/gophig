# Mirror:
⚠️ This is a [mirror](https://git.restartfu.com/restart/git-mirror) of [this repository](https://git.restartfu.com/restart/gophig)
## Getting Started

Gophig may be imported using `go get`:
```
go get git.restartfu.com/restart/gophig.git
```

## Usage

You may create a new `*Gophig`:
```go
type Foo struct{
		foo string `toml:"foo"`
		bar string `toml:"bar"`
}

g := gophig.NewGophig[Foo]("./config.toml", gophig.TOMLMarshaler, os.ModePerm)
```
Then you may use the method `WriteConf(v any)`:
```go
myFooStruct := Foo{foo: "foo", bar: "bar"}

if err := g.WriteConf(myFooStruct); err != nil{
   log.Fatalln(err)
}

// Output file content:
// ./config.toml
/*
   foo = "foo"
   bar = "bar"
*/
```
Or the method `ReadConf[T any]() T`:
```go
// If we assume that the output file content is the same as the example up there:
myFooStruct, err := g.ReadConf[Foo]()
if err != nil {
	log.Fatalln(err)
}

log.Println(foo)

// Output:
/*
   {foo: "foo", bar: "bar"}
*/
```
