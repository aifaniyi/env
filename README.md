# env

A golang library to load environment variables as golang primitives

## Installation
```bash
go get github.com/aifaniyi/env
```

## Usage

Export environment variables in your terminal

```bash
# linux, macos
$ export INT_ENV_VAR=1000

# windows
$ set INT_ENV_VAR=1000
```

```golang
import (
    "fmt"

    "github.com/aifaniyi/env"
)

func main() {
    loadedInt := env.LoadInt("INT_ENV_VAR", 50)
    fmt.Println(loadedInt) // output: 1000
}
```