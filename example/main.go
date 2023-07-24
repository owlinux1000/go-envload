package main

import (
	"fmt"

	"github.com/owlinux1000/go-envload"
)

type Config struct {
}

func main() {
	var cfg struct {
		User     string `env:"USER"`
		ApiToken string `env:"API_TOKEN,required"`
		Timeout  string `env:"TIMEOUT,default=60"`
	}
	if err := envload.Load(&cfg); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", cfg)
}
