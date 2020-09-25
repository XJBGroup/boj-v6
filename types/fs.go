package types

import "github.com/spf13/afero"

type Filesystem = afero.Fs
type FilesystemWrapper = afero.Afero
