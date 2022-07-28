package shx

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type FindFilesOpts struct {
	Ext          []string
	IgnorePrefix []string
}

func FindFiles(root string, opts FindFilesOpts) ([]string, error) {
	var out []string
	walkFn := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !shouldSkipPath(d, path, opts) {
			out = append(out, path)
		}
		return nil
	}
	if err := filepath.WalkDir(root, walkFn); err != nil {
		return nil, err
	}
	return out, nil
}

func shouldSkipPath(d fs.DirEntry, path string, opts FindFilesOpts) bool {
	if d.IsDir() {
		return true
	}

	for _, ext := range opts.Ext {
		if filepath.Ext(path) != ext {
			return true
		}
	}

	for _, prefix := range opts.IgnorePrefix {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}

	return false
}
