package reorder

import (
	"fmt"
	"strconv"
)

// Filer is used to abstract the file actions
type Filer interface {
	GetFiles() []string // GetFiles returns the filepaths in a slice
	Rename(old, new string) error
}

// Run the reorder action using a filer
func Run(f Filer, gap, size int) error {
	fs := f.GetFiles()
	format := createFormatString(size)
	for i, ff := range fs {
		base := string(ff[size+1:])
		new := fmt.Sprintf(format, gap*(i+1), base)
		f.Rename(ff, new)
	}
	return nil
}

func createFormatString(size int) string {
	format := "%0" + strconv.FormatInt(int64(size), 10) + "d_%s"
	return format
}
