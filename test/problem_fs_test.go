package tests

import (
	"github.com/spf13/afero"
	"testing"
)

func TestAbstractFilesystem(t *testing.T) {
	afero.NewOsFs()
}
