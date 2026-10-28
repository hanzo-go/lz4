# LZ4

[![PkgGoDev](https://pkg.go.dev/badge/hanzo.ai/lz4/v4)](https://pkg.go.dev/hanzo.ai/lz4/v4)
[![CI](https://github.com/hanzo-go/lz4/actions/workflows/ci.yml/badge.svg)](https://github.com/hanzo-go/lz4/actions/workflows/ci.yml)

LZ4 compression and decompression written in Go.

## Installation

```bash
go get hanzo.ai/lz4/v4
```

## Example

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"hanzo.ai/lz4/v4"
)

func main() {
	// Compress
	buf := new(bytes.Buffer)
	zw := lz4.NewWriter(buf)
	_, err := io.WriteString(zw, "Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	// Decompress
	zr := lz4.NewReader(buf)
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
```

## Credits

Based on the original implementation by Pierre Curto and contributors at [pierrec/lz4](https://github.com/pierrec/lz4).
