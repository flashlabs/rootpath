package rootpath

import (
	"flag"
	"os"
	"testing"

	"github.com/flashlabs/rootpath/internal"
)

func init() {
	testing.Init()
	flag.Parse()

	if err := os.Chdir(internal.Dir()); err != nil {
		panic(err)
	}
}
