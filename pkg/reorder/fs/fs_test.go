package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFiler_RenameFile(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "filerTest")
	if err != nil {
		t.Fatal("can not create tmpDir")
	}
	fi1, _ := ioutil.TempFile(tmpDir, "001_*")
	fi1.WriteString("some data")
	fi1.Close()
	testFiler := NewFiler(tmpDir, FileMode)
	defer func() {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Println("cannot remove tmpDir", err)
		}
	}()
	type args struct {
		old string
		new string
	}
	tests := []struct {
		name    string
		f       Filer
		args    args
		wantErr bool
	}{
		{
			"FileMode",
			testFiler,
			args{filepath.Base(fi1.Name()), "newName"},
			false,
		},
		{
			"FileMode with not existing file",
			testFiler,
			args{"notexist", "anotherName"},
			true,
		},
		{
			"rename newName",
			testFiler,
			args{"newName", "somethingDifferent"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.Rename(tt.args.old, tt.args.new); (err != nil) != tt.wantErr {
				t.Errorf("Filer.Rename() error = %v, wantErr %v", err, tt.wantErr)
			}
			// do not check the files when error is expected
			if tt.wantErr {
				return
			}
			// Check if new file exists
			if !checkPath(testFiler, tt.args.new) {
				t.Error("new file does not exist", tt.args.new)
			}
			if checkPath(testFiler, tt.args.old) {
				t.Error("old file still exist", tt.args.old)
			}
		})
	}
}

func TestFiler_RenameFolder(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "filerTestFolder")
	if err != nil {
		t.Fatal("can not create tmpDir")
	}
	dir1, _ := ioutil.TempDir(tmpDir, "001_")
	dir11, _ := ioutil.TempDir(dir1, "")
	testFiler := NewFiler(tmpDir, FolderMode)
	defer func() {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Println("cannot remove tmpDir", err)
		}
	}()
	type args struct {
		old string
		new string
	}
	tests := []struct {
		name    string
		f       Filer
		args    args
		wantErr bool
	}{
		{
			"FolderMode",
			testFiler,
			args{filepath.Base(dir1), "newName"},
			false,
		},
		{
			"FolderMode with not existing file",
			testFiler,
			args{"notexist", "anotherName"},
			true,
		},
		{
			"rename newName",
			testFiler,
			args{"newName", "somethingDifferent"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.Rename(tt.args.old, tt.args.new); (err != nil) != tt.wantErr {
				t.Errorf("Filer.Rename() error = %v, wantErr %v", err, tt.wantErr)
			}
			// do not check the files when error is expected
			if tt.wantErr {
				return
			}
			// Check if new folder exists
			if !checkPath(testFiler, filepath.Join(tt.args.new, filepath.Base(dir11))) {
				t.Error("new file does not exist", tt.args.new)
			}
			if checkPath(testFiler, tt.args.old) {
				t.Error("old file still exist", tt.args.old)
			}
		})
	}
}

func checkPath(filer Filer, f string) bool {
	_, err := os.Stat(filepath.Join(filer.fpath, f))
	return !os.IsNotExist(err)

}

func TestFiler_GetElements(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "filerTest")
	if err != nil {
		t.Fatal("can not create tmpDir")
	}
	fi1, _ := ioutil.TempFile(tmpDir, "001_*")
	fi1.Close()
	fi2, _ := ioutil.TempFile(tmpDir, "010_*")
	fi2.Close()
	dir1, _ := ioutil.TempDir(tmpDir, "011_")
	dir2, _ := ioutil.TempDir(tmpDir, "012_")
	defer func() {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Println("cannot remove tmpDir", err)
		}
	}()
	tests := []struct {
		name string
		f    Filer
		want []string
	}{
		{
			"fileMode",
			NewFiler(tmpDir, FileMode),
			[]string{
				filepath.Base(fi1.Name()),
				filepath.Base(fi2.Name()),
			},
		},
		{
			"folder",
			NewFiler(tmpDir, FolderMode),
			[]string{
				filepath.Base(dir1),
				filepath.Base(dir2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.GetElements(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filer.GetElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
