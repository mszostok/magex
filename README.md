# Magefile Extensions

This library provides helper methods to use with [mage](https://magefile.org).

Below is a sample of the type of helpers available. Full examples and documentation is on [godoc](https://godoc.org/github.com/mszostok/magex).

```go
package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	"go.szostok.io/magex/printer"
	"go.szostok.io/magex/shx"
)

func CheckDeadLinks() error {
	printer.Title("Checking dead links in docs...")

	var files []string
	err := filepath.WalkDir("./content", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	}))
	if err != nil {
		return err
	}

	return shx.MustCmdf(`markdown-link-check -q -c .mlc.config.json %s`, strings.Join(files, " ")).
		RunV()
}
```
