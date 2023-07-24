# go-envload

`go-envload` is a tiny library to load enviroment variables to Golang struct.

- **Features**
    - `required` attribute
    - `default=hogehoge` attribute
- **Not supported**
    - Nested struct
    - automatic type cast

## How to use

```go
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
```

