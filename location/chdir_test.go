package location_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flashlabs/rootpath/location"
)

func TestChdir(t *testing.T) {
	assert.NoFileExists(t, "LICENSE")

	if err := location.Chdir(); err != nil {
		panic(err)
	}

	assert.FileExists(t, "LICENSE")
}
