# gozaru

[![Build Status](https://travis-ci.org/subosito/gozaru.svg?branch=master)](https://travis-ci.org/subosito/gozaru)
[![codecov](https://codecov.io/gh/subosito/gozaru/branch/master/graph/badge.svg)](https://codecov.io/gh/subosito/gozaru)

Filename sanitization for Go

## Installation

```
$ go get github.com/subosito/gozaru
```

## Usage

Gozaru basically normalizes, filters and truncates given filename to generates safe and cross-platform filename. For example:

```go
package main

import (
	"fmt"
	"github.com/subosito/gozaru"
)

func main() {
	name := gozaru.Sanitize("  what\\ēver//wëird:user:înput:")
	fmt.Println(name) // => "whatēverwëirduserînput"
}
```

You can add extra room for filename by using `SanitizePad`, see differences here:

```go
// import "strings"

name := strings.Repeat("A", 400)

gozaru.Sanitize(name)
// => resulting filename is 255 characters long

gozaru.SanitizePad(name, 100)
// => resulting filename is 155 characters long
```

## Filenames overview

Best practices for having a safe and cross-platform filenames are:

- Does not contains [ASCII control characters](http://en.wikipedia.org/wiki/ASCII#ASCII_control_characters) (hexadecimal `00` to `1f`)
- Does not contains [Unicode whitespace](http://en.wikipedia.org/wiki/Whitespace_character#Unicode) at the beginning and the end of filename
- Does not contains multiple Unicode whitespaces within the filename
- Does not contains [reserved filenames in Windows](http://msdn.microsoft.com/en-us/library/windows/desktop/aa365247%28v=vs.85%29.aspx)
- Does not contains following characters (according to [wikipedia](http://en.wikipedia.org/wiki/Filename)): `/ \ ? * : | " < >`

## Credits

Gozaru is a Go port of [zaru](https://github.com/madrobby/zaru) by [@madrobby](https://github.com/madrobby). Thanks a lot for him.

