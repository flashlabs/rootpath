package rootpath

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	flagName = "test.v"
	lookup   = "go.mod"
	rootPath = "/"
)

func init() {
	testing.Init()
	flag.Parse()

	err := os.Chdir(dir())
	if err != nil {
		panic(err)
	}
}

func dir() string {
	found := false
	wd, _ := os.Getwd()

	for flag.Lookup(flagName) != nil && wd != rootPath {
		found = hasTarget(wd, lookup)
		if found {
			return wd
		}

		wd = filepath.Dir(wd)
	}

	if !found {
		panic(fmt.Sprintf("can't find the root directory containing %q", lookup))
	}

	return wd
}

func hasTarget(source, target string) bool {
	files, err := os.ReadDir(source)
	if err != nil {
		panic(fmt.Errorf("error while reading a dir %q %w", source, err))
	}

	for _, file := range files {
		if file.Name() == target {
			return true
		}
	}

	return false
}
