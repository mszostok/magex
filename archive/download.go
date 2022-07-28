package archive

import (
	"fmt"
	"os"
	"runtime"

	"github.com/carolynvs/magex/pkg/archive"
	"github.com/carolynvs/magex/pkg/downloads"
)

// Download downloads an archived file to a given bin
func Download(bin string, opts archive.DownloadArchiveOptions) error {
	// determine the appropriate file extension based on the OS, e.g. windows gets .zip, otherwise .tgz
	opts.Ext = opts.ArchiveExtensions[runtime.GOOS]
	if opts.Ext == "" {
		return fmt.Errorf("no archive file extension was specified for the current GOOS (%s)", runtime.GOOS)
	}

	if opts.Hook == nil {
		opts.Hook = archive.ExtractBinaryFromArchiveHook(opts)
	}

	if err := os.MkdirAll(bin, os.ModePerm); err != nil {
		return err
	}

	return downloads.Download(bin, opts.DownloadOptions)
}
