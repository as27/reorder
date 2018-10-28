package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func Test_main(t *testing.T) {
	type flags struct {
		gap     int
		size    int
		minSize int
	}
	tests := []struct {
		name        string
		flags       flags
		inFiles     []string
		inFolders   []string
		wantFiles   []string
		wantFolders []string
	}{
		{
			"default values",
			flags{10, 3, 3},
			[]string{"001_file1", "002_file2"},
			[]string{"011_folder1", "012_folder2"},
			[]string{"010_file1", "020_file2"},
			[]string{"010_folder1", "020_folder2"},
		},
		{
			"longer digits",
			flags{10, 4, 2},
			[]string{"0001_file1", "02_file2", "abc", "1_dd"},
			[]string{"011_folder1", "012_folder2"},
			[]string{"0010_file1", "0020_file2", "abc", "1_dd"},
			[]string{"0010_folder1", "0020_folder2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir, err := ioutil.TempDir("", "reordermain")
			if err != nil {
				t.Error("cannot create tmpDir: ", err)
			}
			defer removeDir(tmpDir)
			createFiles(tmpDir, tt.inFiles)
			checkDir(t, tmpDir, tt.inFiles, true)
			createFolders(tmpDir, tt.inFolders)
			checkDir(t, tmpDir, tt.inFolders, false)
			args = []string{tmpDir}
			*flagGap = tt.flags.gap
			*flagSize = tt.flags.size
			*flagMinSize = tt.flags.minSize
			main()
			checkDir(t, tmpDir, tt.wantFiles, true)
			checkDir(t, tmpDir, tt.wantFolders, false)

		})
	}
}

func createFolders(dir string, ff []string) {
	for _, f := range ff {
		err := os.MkdirAll(filepath.Join(dir, f), 0777)
		checkErr("cannot create folder", err)
	}
}

func createFiles(dir string, ff []string) {
	for _, f := range ff {
		file, err := os.Create(filepath.Join(dir, f))
		checkErr("cannot create file", err)
		file.Close()
	}
}

func checkDir(t *testing.T, dir string, ff []string, checkFiles bool) {
	fmap := make(map[string]bool)
	for _, f := range ff {
		fmap[filepath.Base(f)] = true
	}
	fs, err := ioutil.ReadDir(dir)
	checkErr("cannot read checkFiles", err)
	for _, finfo := range fs {
		if checkFiles && finfo.IsDir() {
			continue
		}
		if !checkFiles && !finfo.IsDir() {
			continue
		}
		_, ok := fmap[filepath.Base(finfo.Name())]
		if !ok {
			t.Errorf("%s should not exist here", finfo.Name())

		}
	}
}

func removeDir(dir string) {
	err := os.RemoveAll(dir)
	checkErr("cannot remove tmpDir", err)
}

func checkErr(s string, err error) {
	if err != nil {
		fmt.Println(s, " :", err)
	}
}
