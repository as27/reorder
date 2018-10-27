package reorder

import (
	"errors"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		f    *testFiler
		gap  int
		size int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		renames []testRename
	}{
		{
			"simple case",
			args{
				makeTestFiler([]string{
					"034_a.txt",
					"035_b.txt",
					"036_c.txt",
				}),
				10,
				3,
			},
			false,
			[]testRename{
				{"034_a.txt", "010_a.txt"},
				{"035_b.txt", "020_b.txt"},
				{"036_c.txt", "030_c.txt"},
			},
		},
		{
			"some other files",
			args{
				makeTestFiler([]string{
					"034_a.txt",
					"abc.txt",
					"1_abc.txt",
					"035_b.txt",
					"1234abc.txt",
					"036_c.txt",
				}),
				10,
				3,
			},
			false,
			[]testRename{
				{"034_a.txt", "010_a.txt"},
				{"035_b.txt", "020_b.txt"},
				{"036_c.txt", "030_c.txt"},
			},
		},
		{
			"longer input digits",
			args{
				makeTestFiler([]string{
					"034_a.txt",
					"035_b.txt",
					"1036_c.txt",
				}),
				10,
				3,
			},
			false,
			[]testRename{
				{"034_a.txt", "010_a.txt"},
				{"035_b.txt", "020_b.txt"},
				{"1036_c.txt", "030_c.txt"},
			},
		},
		{
			"folder",
			args{
				makeTestFiler([]string{
					"034_ab",
					"035_bc",
					"036_cd",
				}),
				10,
				3,
			},
			false,
			[]testRename{
				{"034_ab", "010_ab"},
				{"035_bc", "020_bc"},
				{"036_cd", "030_cd"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(tt.args.f, tt.args.gap, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := tt.args.f.renames
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

func (tf *testFiler) GetElements() []string {
	return tf.files
}

func (tf *testFiler) Rename(old, new string) error {
	if new == "error" {
		return errors.New("error renaming")
	}
	tf.renames = append(tf.renames, testRename{old, new})
	return nil
}

func Test_createFormatString(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"2",
			args{2},
			"%02d_%s",
		},
		{
			"3",
			args{3},
			"%03d_%s",
		},
		{
			"4",
			args{4},
			"%04d_%s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFormatString(tt.args.size); got != tt.want {
				t.Errorf("createFormatString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileBase(t *testing.T) {
	type args struct {
		s    string
		size int
	}
	tests := []struct {
		name     string
		args     args
		wantBase string
		wantOk   bool
	}{
		{
			"simple base",
			args{"000_abc.txt", 3},
			"abc.txt",
			true,
		},
		{
			"simple base",
			args{"00000_abc.txt", 3},
			"abc.txt",
			true,
		},
		{
			"no digits",
			args{"abc.txt", 3},
			"abc.txt",
			false,
		},
		{
			"not enough digits",
			args{"00_abc.txt", 3},
			"00_abc.txt",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBase, gotOk := fileBase(tt.args.s, tt.args.size)
			if gotBase != tt.wantBase {
				t.Errorf("fileBase() gotBase = %v, want %v", gotBase, tt.wantBase)
			}
			if gotOk != tt.wantOk {
				t.Errorf("fileBase() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
