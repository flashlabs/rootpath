package location

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/flashlabs/rootpath/internal"
)

func Chdir() error {
	testing.Init()
	flag.Parse()

	if err := os.Chdir(internal.Dir()); err != nil {
		return fmt.Errorf("error while executing the os.Chdir: %w", err)
	}

	return nil
}
