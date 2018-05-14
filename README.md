# countries-emoji
Emoji characters for all countries.

# Getting Started

## Installing

Install Go and run `go get`:
```bash
$ go get -u github.com/TimothyYe/countries-emoji
```

## Usage

### Function

```go
func GetEmojiFlag(countryCode string) string 
```

Parameter `countryCode` is an [ISO 3166-1 alpha-2 code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)

### Example

```go
package main

import (
        "fmt"

        ce "github.com/TimothyYe/countries-emoji"
)

func main() {
        fmt.Println(ce.GetEmojiFlag("CN"))
}
```

output is:
```bash
🇨🇳
```

