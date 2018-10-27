package reorder

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		f Filer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		renames []testRename
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func makeTestFiler(f []string) Filer {
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
