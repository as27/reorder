package reorder

import (
	"errors"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	tests := []struct {
		name    string
		f       *testFiler
		wantErr bool
		renames []testRename
	}{
		{
			"simple case",
			makeTestFiler([]string{
				"034_a.txt",
				"035_b.txt",
				"036_c.txt",
			}),
			false,
			[]testRename{
				{"034_a.txt", "010_a.txt"},
				{"035_b.txt", "020_b.txt"},
				{"036_c.txt", "030_c.txt"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(tt.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := tt.f.renames
			if !reflect.DeepEqual(got, tt.renames) {
				t.Errorf("Renames = %v, want %v", got, tt.renames)
			}
		})
	}
}

func makeTestFiler(f []string) *testFiler {
	return &testFiler{
		files: f,
	}
}

type testFiler struct {
	files   []string
	renames []testRename
}

type testRename struct {
	old, new string
}

func (tf *testFiler) GetFiles() []string {
	return tf.files
}

func (tf *testFiler) Rename(old, new string) error {
	if new == "error" {
		return errors.New("error renaming")
	}
	tf.renames = append(tf.renames, testRename{old, new})
	return nil
}
