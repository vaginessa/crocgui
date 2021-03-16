package main

import (
	"net/url"
	"strings"
)

func init() {
	// TODO: android probably has a way to find this
	DEFAULT_DOWNLOAD_DIR = "/storage/emulated/0/Download"
}

func fixpath(fpath string) string {
	if strings.Contains(fpath, "%3A") {
		fpath, _ = url.PathUnescape(fpath)
		colIdx := strings.Index(fpath, ":")
		fpath = "/storage/emulated/0/" + fpath[colIdx+1:]
	}

	return fpath
}
