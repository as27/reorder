// Package fs implements the file system for the reorder interface
package fs

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Filer is the reorder.Filer implementation
type Filer struct {
	fpath string
	mode  Mode
}

// Mode for the Filer
type Mode int

// Different modes how the filer can be used
const (
	FileMode   Mode = iota // use just files
	FolderMode             // use just folders
)

// NewFiler creates a new Filer
func NewFiler(fpath string, mode Mode) Filer {
	return Filer{
		fpath: fpath,
		mode:  mode,
	}
}

// GetElements returns files or folder, depends on the
// mode of the filer
func (f Filer) GetElements() []string {
	var elements []string
	fileInfos, err := ioutil.ReadDir(f.fpath)
	if err != nil {
		log.Println("error reading dir:", err)
		return elements
	}
	for _, fi := range fileInfos {
		switch f.mode {
		case FileMode:
			if fi.IsDir() {
				continue
			}
		case FolderMode:
			if !fi.IsDir() {
				continue
			}
		}
		elements = append(elements, filepath.Base(fi.Name()))
	}
	return elements
}

// Rename let you rename a folder or file
func (f Filer) Rename(old, new string) error {
	return os.Rename(
		filepath.Join(f.fpath, filepath.Base(old)),
		filepath.Join(f.fpath, filepath.Base(new)),
	)
}
