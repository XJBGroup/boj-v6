package unittest

import (
	"encoding/gob"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_statics"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
	"io"
	"os"
)

func serialize(w io.Writer, data *GoDynamicTestData) error {
	return gob.NewEncoder(w).Encode(data)
}

func deserialize(r io.Reader, data *GoDynamicTestData) error {
	return gob.NewDecoder(r).Decode(data)
}

func loadCache(filepath string) (d *GoDynamicTestData, err error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	d = new(GoDynamicTestData)
	return d, deserialize(f, d)
}

func saveCache(filepath string, d *GoDynamicTestData) (err error) {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	return serialize(f, d)
}

func init() {
	gob.Register(new(GoDynamicTestData))
	gob.Register(unittest_statics.StringMap{})
	gob.Register(unittest_types.DataBody{})
}
