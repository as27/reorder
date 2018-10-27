package reorder

// Filer is used to abstract the file actions
type Filer interface {
	GetFiles() []string // GetFiles returns the filepaths in a slice
	Rename(old, new string) error
}

// Run the reorder action using a filer
func Run(f Filer, gap, size int) error {
	fs := f.GetFiles()
	for _, ff := range fs {
		f.Rename(ff, ff)
	}
	return nil
}
