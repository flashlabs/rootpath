package manual_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flashlabs/rootpath/manual"
)

func TestChdir(t *testing.T) {
	assert.NoFileExists(t, "LICENSE")

	if err := manual.Chdir(); err != nil {
		panic(err)
	}

	assert.FileExists(t, "LICENSE")
}
