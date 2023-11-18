package tests_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/flashlabs/rootpath"
)

func TestRootpath(t *testing.T) {
	assert.FileExists(t, "LICENSE")
}
