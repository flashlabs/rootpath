package test_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/flashlabs/rootpath"
)

func TestRootpathBlankImport(t *testing.T) {
	assert.FileExists(t, "LICENSE")
}
