# norma

[![Build Status](https://drone.io/github.com/subosito/norma/status.png)](https://drone.io/github.com/subosito/norma/latest)
[![Coverage Status](https://coveralls.io/repos/subosito/norma/badge.png?branch=master)](https://coveralls.io/r/subosito/norma?branch=master)

Filename sanitization for Go

## Installation

```
$ go get github.com/subosito/norma
```

## Usage

Norma basically normalizes, filters and truncates given filename to generates safe and cross platform filename. For example:

```go
package main

import (
	"fmt"
	"github.com/subosito/norma"
)

func main() {
	name := norma.Sanitize("  what\\ēver//wëird:user:înput:")
	fmt.Println(name) // => "whatēverwëirduserînput"
}
```

You can add extra room for filename by using `SanitizePad`, see differences here:

```go
// import "strings"

name := strings.Repeat("A", 400)

norma.Sanitize(name)
// => resulting filename is 255 characters long

norma.SanitizePad(name, 100)
// => resulting filename is 155 characters long
```

## Credits

Norma is a Go port of [zaru](https://github.com/madrobby/zaru) by [@madrobby](https://github.com/madrobby). Thanks a lot for him.

