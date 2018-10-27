package reorder

import (
	"fmt"
	"path/filepath"
	"regexp"
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
	i := 1
	for _, ff := range fs {
		base, ok := fileBase(ff, size)
		if !ok {
			// just ignore files, which do not match the
			// digit logic
			continue
		}
		new := fmt.Sprintf(format, gap*i, base)
		f.Rename(ff, new)
		i++
	}
	return nil
}

func createFormatString(size int) string {
	format := "%0" + strconv.FormatInt(int64(size), 10) + "d_%s"
	return format
}

// fileBase removes the digits at the beginning of the file name
// the size defines the minimum of the digits
func fileBase(s string, size int) (base string, ok bool) {
	s = filepath.Base(s)
	re := regexp.MustCompile(fmt.Sprintf("\\d{%d,}_", size))
	i := re.FindStringIndex(s)
	if i == nil {
		return s, false
	}
	return s[i[1]:], true
}
