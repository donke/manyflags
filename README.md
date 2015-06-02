# manyflags
split many flags!

## Usage
```go
package main

import (
    "fmt"
    "os"

    "github.com/donke/manyflags"
)

var (
    a = flag.Bool("a", false, "")
    b = flag.Bool("b", false, "")
    c = flag.Bool("c", false, "")
)

func main() {
    os.Args = append(os.Args, "-abc")

    manyflags.OverwriteArgs()
    flag.Parse()

    fmt.Println(*a, *b, *c) // true true true
}
```

## Installation
```
$ go get github.com/donke/manyflags
```

## License

MIT

## Author

Ken Kudo (aka.kudoken@gmail.com)
