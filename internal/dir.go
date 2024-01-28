package internal

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Dir() string {
	found := false
	wd, _ := os.Getwd()

	for flag.Lookup(FlagName) != nil && wd != RootPath {
		found = hasTarget(wd, Lookup)
		if found {
			return wd
		}

		wd = filepath.Dir(wd)
	}

	if !found {
		panic(fmt.Sprintf("can't find the root directory containing %q", Lookup))
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
