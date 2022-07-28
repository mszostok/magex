package deps

import (
	"github.com/carolynvs/magex/pkg"
	"github.com/carolynvs/magex/pkg/archive"
	"github.com/carolynvs/magex/pkg/downloads"

	archivex "go.szostok.io/magex/archive"
	"go.szostok.io/magex/printer"
)

func EnsureGolangciLint(bin, version string) error {
	printer.Titlef("Ensure golangci-lint in ./bin (ver: %s)", version)

	found, err := pkg.IsCommandAvailable("./bin/golangci-lint", "version", version)
	if err != nil {
		return err
	}
	if found {
		return nil
	}

	opts := archive.DownloadArchiveOptions{
		DownloadOptions: downloads.DownloadOptions{
			UrlTemplate: "https://github.com/golangci/golangci-lint/releases/download/v{{.VERSION}}/golangci-lint-{{.VERSION}}-{{.GOOS}}-{{.GOARCH}}{{.EXT}}",
			Name:        "golangci-lint",
			Version:     "1.47.2",
		},
		ArchiveExtensions: map[string]string{
			"linux":   ".tar.gz",
			"darwin":  ".tar.gz",
			"windows": ".zip",
		},
		TargetFileTemplate: "golangci-lint-{{.VERSION}}-{{.GOOS}}-{{.GOARCH}}/golangci-lint{{.EXT}}",
	}
	return archivex.Download(bin, opts)
}

func EnsureMdox(bin, version string) error {
	return pkg.EnsurePackageWith(pkg.EnsurePackageOptions{
		Name:           "github.com/bwplotka/mdox",
		VersionCommand: "--version",
		Destination:    "bin",
		DefaultVersion: version,
	})
}
