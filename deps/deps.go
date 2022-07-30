package deps

import (
	"path/filepath"

	"github.com/carolynvs/magex/pkg"
	"github.com/carolynvs/magex/pkg/archive"
	"github.com/carolynvs/magex/pkg/downloads"

	archivex "go.szostok.io/magex/archive"
	"go.szostok.io/magex/printer"
)

func EnsureGolangciLint(bin, version string) error {
	printer.Titlef("Ensure golangci-lint in %s (ver: %s)", bin, version)

	found, err := pkg.IsCommandAvailable(filepath.Join(bin, "golangci-lint"), "version", version)
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
			Version:     version,
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

func EnsureMuffet(bin, version string) error {
	printer.Titlef("Ensure muffet in %s (ver: %s)", bin, version)

	found, err := pkg.IsCommandAvailable(filepath.Join(bin, "muffet"), "--version", version)
	if err != nil {
		return err
	}
	if found {
		return nil
	}

	opts := archive.DownloadArchiveOptions{
		DownloadOptions: downloads.DownloadOptions{
			UrlTemplate: "https://github.com/raviqqe/muffet/releases/download/v{{.VERSION}}/muffet_{{.VERSION}}_{{.GOOS}}_{{.GOARCH}}{{.EXT}}",
			Name:        "muffet",
			Version:     version,
			OsReplacement: map[string]string{
				"darwin":  "Darwin",
				"linux":   "Linux",
				"windows": "Windows",
			},
			ArchReplacement: map[string]string{
				"amd64": "x86_64",
				"386":   "i386",
			},
		},
		ArchiveExtensions: map[string]string{
			"linux":   ".tar.gz",
			"darwin":  ".tar.gz",
			"windows": ".tar.gz",
		},
		TargetFileTemplate: "muffet{{.EXT}}",
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
