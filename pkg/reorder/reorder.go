/*
Package reorder is for reordering a set of elements. This could
be files or folders. The elements have to start with digits.
The size of the reorder defines the minum of the digits of the
element.
*/
package reorder

import (
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

// Filer is used to abstract the file actions
type Filer interface {
	GetElements() []string // GetElements returns files or folder
	Rename(old, new string) error
}

// Run the reorder action using a filer
func Run(f Filer, gap, size, minSize int) error {
	fs := f.GetElements()
	sort.Sort(elements(fs))
	format := createFormatString(size)
	i := 1
	for _, ff := range fs {
		base, ok := fileBase(ff, minSize)
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

func fileNumber(s string) int {
	s = filepath.Base(s)
	re := regexp.MustCompile("\\d{1,}")
	i := re.FindStringIndex(s)
	if i == nil {
		return 0
	}
	number, err := strconv.Atoi(s[:i[1]])
	if err != nil {
		return 0
	}
	return number
}

// elements is used to implement the sorter interface
type elements []string

func (e elements) Len() int      { return len(e) }
func (e elements) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e elements) Less(i, j int) bool {
	return fileNumber(e[i]) < fileNumber(e[j])
}
