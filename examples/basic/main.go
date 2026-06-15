package main

import (
	"log"
	"os"

	"github.com/restartfu/gophig"
	"github.com/restartfu/gophig/codecs"
)

type Config struct {
	Name string `toml:"name"`
	Port int    `toml:"port"`
}

func main() {
	conf := gophig.NewGophig[Config]("config.toml", codecs.TOMLMarshaler{}, 0o644)

	if err := conf.SaveConf(Config{Name: "gophig", Port: 8080}); err != nil {
		log.Fatal(err)
	}
	defer os.Remove("config.toml")

	loaded, err := conf.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", loaded)
}
